import React, { useEffect, useState } from 'react'
import { useOutletContext, useParams } from 'react-router'
import { NotificationManager } from 'react-notifications'
import formatDate, { calculateReadingTime } from '../utils/Utility'
import DOMPurify from 'dompurify'
import hljs from "highlight.js"
import 'highlight.js/styles/monokai.css';

const Blog = () => {
    const { publicAxiosInstance } = useOutletContext()
    let { blog_url } = useParams();
    const [blog, setBlog] = useState({});
    const [readingTime, setReadingTime] = useState(0)
    const [htmlContent, setHtmlContent] = useState('')

    const getBlog = () => {
        publicAxiosInstance.get(`/blog/${blog_url}`)
            .then(response => {
                setBlog(response.data)
                setReadingTime(calculateReadingTime(response.data.content.split(" ").length))
                const sanitizedHtml = DOMPurify.sanitize(response.data.content)
                setHtmlContent(sanitizedHtml)

            })
            .catch(err => {
                NotificationManager.error(err)
            })
    }

    useEffect(() => {
        getBlog();
    }, [])

    useEffect(() => {
        document.querySelectorAll("pre").forEach((block) => {
            hljs.highlightElement(block) // highlights code on code blocks on website

            block.addEventListener("dblclick", (event) => {
                const codeText = block.innerText;
                navigator.clipboard.writeText(codeText)
                    .then(() => {
                        NotificationManager.success("Code copied to clipboard")
                    }).catch(err => {
                        NotificationManager.error("Failed to cpy code.")
                    });
            });
        })
    }, [htmlContent])

    return (
        <div className="mt-[132.5px] mb-10  blogSection dark:text-white p-3 container m-auto ">
            <div className='mx-5'>
                <h1 className='class="text-center text-xl md:text-3xl justify-center lg:text-4xl font-semibold text-gray-800 dark:text-white mb-1 mt-[20px] flex"'>{blog.title}</h1>
                <div className='class="flex items-center mb-6 flex-col md:flex-row md:justify-center"'>{blog.author} &nbsp; &middot;
                    <span className='text-slate-400'>
                        &nbsp; {readingTime} min read </span> </div>
                <div className="cont leading-relaxed text-dark dark:text-gray-100" dangerouslySetInnerHTML={{ __html: htmlContent }} />
            </div>
        </div>
    )
}

export default Blog