import React, { useState } from 'react';
import { useNavigate, useOutletContext } from 'react-router';
import TextEditor from '../utils/TextEditor';
import { NotificationManager } from 'react-notifications';
const AdminAddBlog = () => {
    const navigate = useNavigate()
    const { privateAxiosInstance } = useOutletContext();
    const [newBlog, setNewBlog] = useState({
        title: "",
        content: "",
        author: "",
        tags: [],
        url: "",
        admin_mode: 0,
        is_admin: true,
        is_hidden: false,
    });

    const handleAddFormSubmit = async (event) => {
        event.preventDefault();
        const config = {
            headers: {
                'Content-Type': 'application/json',
                // 'Authorization': 'Bearer YOUR_ACCESS_TOKEN'
            }
        };
        try {
            // console.log("newBlog: ", newBlog)
            const response = await privateAxiosInstance.post('/blog', newBlog, config);
            // console.log("form submit....", response);
            navigate(`/blogs/${newBlog.url}`)
            NotificationManager.success("Blog Added Successfully!")
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

    const handleChange = (event) => {
        const { name, value } = event.target;
        if (name === 'tags') {
            setNewBlog({
                ...newBlog,
                [name]: value.split(',').map(tag => tag.trim())
            });
        } else {
            setNewBlog({
                ...newBlog,
                [name]: value
            });
        }
    };

    const handleContentChange = (content) => {
        setNewBlog({
            ...newBlog,
            content
        });
    }

    return (
        <>
            <div className="p-6">
                <h1 className='text-3xl text-center'>Write A Blog</h1>
                <form onSubmit={handleAddFormSubmit} className='flex flex-col space-y-3'>
                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.title} type="text" id="title" name='title' onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Content</label>
                    <TextEditor editorHtml={newBlog.content} setEditorHtml={handleContentChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.tags.join(', ')} type="text" name="tags" id="tags" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.url} type="text" name="url" id="url" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newBlog.author} type="text" name="author" id="author" onChange={handleChange} />

                    <button type='submit' className='w-40 py-3 px-5 dark:bg-slate-800 bg-gray-500 border font-bold hover:dark:bg-slate-900'>Add Blog</button>
                </form>
            </div>
        </>
    )
}

export default AdminAddBlog;
