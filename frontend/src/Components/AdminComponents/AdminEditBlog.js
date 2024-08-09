import React, { useEffect, useState } from 'react'
import { useNavigate, useOutletContext, useParams } from 'react-router'
import TextEditor from '../utils/TextEditor';
import { NotificationManager } from 'react-notifications';
const AdminEditBlog = () => {
    const navigate = useNavigate()
    const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
    const [editorHtml, setEditorHtml] = useState('');
    const [oldBlog, setOldBlog] = useState({})
    const [newBlog, setNewBlog] = useState(oldBlog)

    let { blog_url } = useParams()
    let blogHeading = blog_url.split("-").map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(" ")

    const getOldBlogByUrl = async (url) => {
        try {
            const response = await publicAxiosInstance.get(`/blog/${url}`);
            // console.log(response.data);
            setOldBlog(response.data);
            setNewBlog(response.data);
        } catch (err) {
            if (err.response) {
                NotificationManager.error(`Data: ${err.response.data}\nStatus: ${err.response.status}\nHeaders: ${err.response.headers}`)
            } else if (err.request) {
                NotificationManager.error(`${err.message}`)
            } else {
                NotificationManager.error(`${err.message}`)
            }
        }
    }

    const handleEditFormSubmit = async (event) => {
        event.preventDefault();
        const config = {
            headers: {
                'Content-Type': 'application/json',
                // 'Authorization': 'Bearer YOUR_ACCESS_TOKEN'
            }
        };
        try {
            const response = await privateAxiosInstance.put(`/blog/${blog_url}`, newBlog, config);
            // console.log("form submit....", response.data.updated_fields.url);
            NotificationManager.success("Blog Updated Successfully!")
            navigate(`/blogs/${response.data.updated_fields.url}`)
        } catch (err) {
            if (err.response) {
                NotificationManager.error(`Data: ${err.response.data}\nStatus: ${err.response.status}\nHeaders: ${err.response.headers}`)
            } else if (err.request) {
                NotificationManager.error(`${err.message}`)
            } else {
                NotificationManager.error(`${err.message}`)
            }
        }
    }

    useEffect(() => {
        getOldBlogByUrl(blog_url);
    }, [blog_url]);

    const handleChange = (event) => {
        const { name, value } = event.target;
        setNewBlog({
            ...newBlog,
            [name]: value
        });
    };

    const handleContentChange = (content) => {
        setNewBlog({
            ...newBlog,
            content
        });
        // console.log(content);
    }

    return (
        <>
            <div className="p-6">
                <h1 className='text-3xl text-center'>Update {blogHeading}</h1>
                <form onSubmit={handleEditFormSubmit} className='flex flex-col space-y-3'>
                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.title} type="text" id="title" name='title' onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Content</label>
                    <TextEditor editorHtml={newBlog.content} setEditorHtml={handleContentChange} />

                    <label className='px-2 font-bold  hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.tags} type="text" name="tags" id="tags" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none  h-12' value={newBlog.url} type="text" name="url" id="url" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.author} type="text" name="author" id="author" onChange={handleChange} />

                    <button type='submit' className='w-40 py-3 px-5 dark:bg-slate-800  bg-gray-500 border font-bold hover:dark:bg-slate-900 '>Update</button>
                </form>
            </div>
        </>
    )
}

export default AdminEditBlog;
