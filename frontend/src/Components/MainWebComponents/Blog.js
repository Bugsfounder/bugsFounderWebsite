import React, { useEffect, useState } from 'react'
import { useOutletContext, useParams } from 'react-router'
import { NotificationManager } from 'react-notifications'
import formatDate, { calculateReadingTime } from '../utils/Utility'
import DOMPurify from 'dompurify'

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

    return (

        <div className="mt-[132.5px] blogSection dark:text-white p-3">
            <h1 className='class="text-center text-xl md:text-3xl justify-center lg:text-4xl font-semibold text-gray-800 dark:text-white mb-1 mt-[20px] flex"'>{blog.title}</h1>
            <div className='class="flex items-center mb-6 flex-col md:flex-row md:justify-center"'>{blog.author} &nbsp; &middot;  &nbsp;
                <span className='text-slate-400'>
                    {/* {formatDate(blog.updated_at)} */}
                    {blog.updated_at}
                    &nbsp; &middot;  &nbsp; {readingTime} min read </span> </div>
            <div className="cont leading-relaxed text-dark dark:text-gray-100" dangerouslySetInnerHTML={{ __html: htmlContent }} />
        </div>
    )
}

export default Blog