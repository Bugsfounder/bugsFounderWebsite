import React, { useEffect, useState } from 'react'
import { useOutletContext, useParams } from 'react-router'

const AdminEditBlog = () => {
    const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
    const [oldBlog, setOldBlog] = useState({})
    const [newBlog, setNewBlog] = useState(oldBlog)

    let { blog_url } = useParams()
    let blogHeading = blog_url.split("-").map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(" ")

    const getOldBlogByUrl = async (url) => {
        const response = await publicAxiosInstance.get(`/blog/${blog_url}`)
        console.log(response.data)
        setOldBlog(response.data)
        setNewBlog(response.data)
    }
    const handleEditFormSubmit = async (event) => {
        event.preventDefault()
        // const blog_url = await 
        console.log("form submit....")
    }

    useEffect(() => {
        getOldBlogByUrl(blog_url)
    }, [])

    const handleChange = (event) => {
        const { name, value } = event.target
        setNewBlog({
            ...newBlog,
            [name]: value
        });
    };

    return (
        <>
            <div className="p-6">
                <h1 className='text-3xl text-center'>Update {blogHeading}</h1>
                <form onSubmit={handleEditFormSubmit} className='flex flex-col space-y-3 text-lg'>
                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                    <input className='dark:bg-slate-600 p-2 rounded-[10px] h-12' value={newBlog.title} placeholder='Add New   if want to edit otherwise leave it blank' type="text" id="title" name='title' onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Content</label>
                    <textarea value={newBlog.content} type="text" name="content" id="content" />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                    <input className='dark:bg-slate-600 p-2 rounded-[10px] h-12' value={newBlog.tags} placeholder='Add New   if want to edit otherwise leave it blank' type="tags" name="tags" id="tags" />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                    <input className='dark:bg-slate-600 p-2 rounded-[10px] h-12' value={newBlog.url} placeholder='Add New   if want to edit otherwise leave it blank' type="text" name="url" id="url" />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                    <input className='dark:bg-slate-600 p-2 rounded-[10px] h-12' value={newBlog.author} placeholder='Add New   if want to edit otherwise leave it blank' type="text" name="author" id="author" />
                </form>
            </div>
        </>
    )
}

export default AdminEditBlog