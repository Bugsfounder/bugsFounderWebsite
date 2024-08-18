import React, { useEffect, useState } from 'react'
import { useNavigate, useOutletContext, useParams } from 'react-router'
import { NotificationManager } from 'react-notifications';
import { Link } from 'react-router-dom';
import { PlusCircleIcon } from '@heroicons/react/24/solid';

const AdminEditTutorial = () => {
    const navigate = useNavigate()
    const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
    const [oldTutorial, setOldTutorial] = useState({})
    const [newTutorial, setNewTutorial] = useState(oldTutorial)

    let { tutorial_url } = useParams()
    let TutorialHeading = tutorial_url.split("-").map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(" ")

    const getOldTutorialByUrl = async (url) => {
        try {
            const response = await publicAxiosInstance.get(`/tutorial/${url}`);
            const tutorialData = response.data;
            console.log("old: ", response.data);

            // Ensure sub_tutorials is initialized as an array
            if (tutorialData.sub_tutorials == null) {
                tutorialData.sub_tutorials = [];
            }
            setOldTutorial(tutorialData);
            setNewTutorial(tutorialData);
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
            const response = await privateAxiosInstance.put(`/tutorial/${tutorial_url}`, newTutorial, config);
            const tutorialData = response.data;

            // Ensure sub_tutorials is initialized as an array
            if (!tutorialData.sub_tutorials) {
                tutorialData.sub_tutorials = [];
            }

            // console.log("form submit....", response, newTutorial.sub_tutorials);
            NotificationManager.success("Tutorial Updated Successfully!")
            navigate(`/tutorials/${response.data.updated_fields.url}`)
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
        getOldTutorialByUrl(tutorial_url);
    }, [tutorial_url]);

    const handleChange = (event) => {
        const { name, value } = event.target;
        setNewTutorial({
            ...newTutorial,
            [name]: value
        });
    };

    return (
        <>
            <div className="p-6">
                <h1 className='text-3xl text-center'>Update {TutorialHeading}</h1>
                <form onSubmit={handleEditFormSubmit} className='flex flex-col space-y-3'>
                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="title">Title</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12 rounded-[10px]' value={newTutorial.title} type="text" id="title" name='title' onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="content">Description</label>
                    <textarea name='description' id='description' className='w-full h-40 p-4 rounded-[10px]  ark:bg-slate-900 border outline-none dark:bg-slate-900' rows="5" value={newTutorial.description} onChange={handleChange} />

                    <label className='px-2 font-bold  hover:text-slate-400 cursor-pointer' htmlFor="tags">Tags</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12 rounded-[10px]' value={newTutorial.tags} type="text" name="tags" id="tags" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="url">Url</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12 rounded-[10px]' value={newTutorial.url} type="text" name="url" id="url" onChange={handleChange} />

                    <label className='px-2 font-bold hover:text-slate-400 cursor-pointer' htmlFor="author">Author</label>
                    <input required className='dark:bg-slate-900 border p-2 outline-none h-12 rounded-[10px]' value={newTutorial.author} type="text" name="author" id="author" onChange={handleChange} />

                    <div className='px-4 flex items-center space-x-10'>
                        <Link to={`/kubari/admin/tutorials/edit/sub_tutorial/${tutorial_url}`}>
                            <span>
                                SubTutorials length: &nbsp;
                                {newTutorial.sub_tutorials?.length || 0}
                            </span>
                        </Link>
                        <Link to={`/kubari/admin/tutorials/edit/sub_tutorial/add/${newTutorial.url}`} title='Add a sub tutorial'>
                            <PlusCircleIcon className="size-10 text-slate-600 hover:text-slate-700 cursor-pointer " />
                        </Link>
                    </div>

                    <button type='submit' className='w-40 py-3 px-5 rounded-[10px] dark:bg-slate-800  bg-gray-500 border font-bold hover:dark:bg-slate-900'>Update</button>
                </form>
            </div>
        </>
    )
}

export default AdminEditTutorial