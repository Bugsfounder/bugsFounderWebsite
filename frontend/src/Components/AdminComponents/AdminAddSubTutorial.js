import React, { useState } from 'react';
import { useNavigate, useOutletContext, useParams } from 'react-router';
import TextEditor from '../utils/TextEditor';
import { NotificationManager } from 'react-notifications';

const AdminAddSubTutorial = () => {
    const navigate = useNavigate()
    const { privateAxiosInstance } = useOutletContext();
    const [newSubTutorial, setNewSubTutorial] = useState({
        title: "",
        content: "",
        author: "",
        tags: [],
        url: "",
        admin_mode: 0,
        is_admin: true,
        is_hidden: false,
    });
    let { tutorial_url } = useParams()

    const handleAddFormSubmit = async (event) => {
        event.preventDefault();

        // Ensure sub_tutorials is an array
        if (!Array.isArray(newSubTutorial.sub_tutorials)) {
            newSubTutorial.sub_tutorials = [];
        }

        const config = {
            headers: {
                'Content-Type': 'application/json',
                // 'Authorization': 'Bearer YOUR_ACCESS_TOKEN'
            }
        };
        try {
            console.log("newSubTutorial: ", newSubTutorial)
            const response = await privateAxiosInstance.post(`/tutorial/${tutorial_url}`, newSubTutorial, config);
            console.log("form submit....", response);
            navigate(`/tutorials/${tutorial_url}/${newSubTutorial.url}`)
            NotificationManager.success("Tutorial Added Successfully!")
        } catch (err) {
            if (err.response) {
                console.error('Error Data:', err.response.data);
                NotificationManager.error(`Error: ${JSON.stringify(err.response.data)}\nStatus: ${err.response.status}`);
            } else if (err.request) {
                console.error('Request Error:', err.request);
                NotificationManager.error(`${err.message}`);
            } else {
                console.error('Error:', err.message);
                NotificationManager.error(`${err.message}`);
            }
        }
    }

    const handleChange = (event) => {
        const { name, value } = event.target;
        if (name === 'tags') {
            setNewSubTutorial({
                ...newSubTutorial,
                [name]: value.split(',').map(tag => tag.trim())
            });
        } else {
            setNewSubTutorial({
                ...newSubTutorial,
                [name]: value
            });
        }
    };

    const handleContentChange = (content) => {
        setNewSubTutorial({
            ...newSubTutorial,
            content
        });
        // console.log(content);
    }

    return (
        <>
            <div className="p-6">
                <h1 className='text-3xl text-center'>Write A Sub Tutorial</h1>
                <form onSubmit={handleAddFormSubmit} className='flex flex-col space-y-3'>
                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newSubTutorial.title} type="text" id="title" name='title' onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Content</label>
                    <TextEditor editorHtml={newSubTutorial.content} setEditorHtml={handleContentChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newSubTutorial.tags.join(', ')} type="text" name="tags" id="tags" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newSubTutorial.url} type="text" name="url" id="url" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12' value={newSubTutorial.author} type="text" name="author" id="author" onChange={handleChange} />

                    <button type='submit' className='w-40 py-3 px-5 dark:bg-slate-800 bg-gray-500 border font-bold hover:dark:bg-slate-900'>Add Tutorial</button>
                </form>
            </div>
        </>
    )
}

export default AdminAddSubTutorial