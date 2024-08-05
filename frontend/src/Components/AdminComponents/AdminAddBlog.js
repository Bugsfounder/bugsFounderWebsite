import React, { useEffect, useState } from 'react'
import { useOutletContext } from 'react-router';
import TextEditor from '../utils/TextEditor';

const AdminAddBlog = () => {
    const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
    const [blog, setBlog] = useState(
        {
            "title": "",
            "content": "",
            "author": "",
            "tags": "",
            "url": "",
            // TODO: these three fields has to be automatically filled by logged in admin
            "admin_mode": 0,
            "is_admin": true,
            "is_hidden": false,
        }
    )

    // const handleAddFormSubmit = (event) => {
    //     event.preventDefault();
    //     const requestOptions = {
    //         method: "POST",
    //         headers: {
    //             'Content-Type': 'application/json',
    //             // 'Authorization': 'Bearer YOUR_ACCESS_TOKEN'
    //         },
    //         body: JSON.stringify(blog)
    //     };
    //     // privateAxiosInstance.post(`/blog`, blog, config)
    //     //     .then(response => {
    //     //         console.log(response);
    //     //     }).catch(err => (
    //     //         console.log("getting error:", err)
    //     //     ))

    //     fetch('http://localhost:8080/api/private/admin/blog', requestOptions)
    //         .then(response => response.json())
    //         .then(data => console.log(data))
    //         .catch(err => console.log("error: ", err))
    // }

    const handleAddFormSubmit = (event) => {
        event.preventDefault();
        privateAxiosInstance.post('/api/private/admin/blog', blog)
            .then(response => {
                console.log(response.data);
            })
            .catch(err => console.log("Error:", err));
    };

    const handleChange = (event) => {
        const { name, value } = event.target;
        setBlog({
            ...blog,
            [name]: value
        });
    };

    const handleContentChange = (content) => {
        setBlog({
            ...blog,
            content
        });
        console.log(content);
    }
    return (
        <div className="p-6">
            <h1 className='text-3xl text-center'>Write A Blog</h1>
            <form onSubmit={handleAddFormSubmit} className='flex flex-col space-y-3'>
                <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={blog.title} type="text" id="title" name='title' onChange={handleChange} />

                <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Content</label>
                <TextEditor editorHtml={blog.content} setEditorHtml={handleContentChange} />

                <label className='px-2 font-bold  hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={blog.tags} type="text" name="tags" id="tags" onChange={handleChange} />

                <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                <input required className='dark:bg-slate-900 border p-2 outline-none  h-12' value={blog.url} type="text" name="url" id="url" onChange={handleChange} />

                <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={blog.author} type="text" name="author" id="author" onChange={handleChange} />

                <button type='submit' className='w-40 py-3 px-5 dark:bg-slate-800  bg-gray-500 border font-bold hover:dark:bg-slate-900 '>Update</button>
            </form>
        </div>
    )
}

export default AdminAddBlog