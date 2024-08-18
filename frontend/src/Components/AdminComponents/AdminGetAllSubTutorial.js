import React, { useEffect, useState } from 'react';
import { Link, useOutletContext, useParams } from 'react-router-dom';
import { PencilSquareIcon, TrashIcon, PlusCircleIcon } from '@heroicons/react/24/solid';
import InfiniteScroll from 'react-infinite-scroll-component';
import ConfirmationModal from '../utils/ConfigmationModal';
import { NotificationManager } from 'react-notifications';
import formatDate from '../utils/Utility';

const AdminGetAllSubTutorial = () => {
    const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
    const [tutorials, setTutorials] = useState([]);
    const [page, setPage] = useState(1);
    const [hasMore, setHasMore] = useState(true);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [tutorialToDelete, setTutorialToDelete] = useState(null);
    const { tutorial_url } = useParams()

    const getTutorials = () => {
        // publicAxiosInstance.get(`/tutorial/${tutorial_url}?page=${page}&limit=6`)
        publicAxiosInstance.get(`/tutorial/${tutorial_url}`)
            // publicAxiosInstance.get(`/tutorial`)
            .then(response => {
                console.log(response);
                const subTutorials = response.data.sub_tutorials
                if (subTutorials && subTutorials.length > 0) {
                    setTutorials(prevTutorials => {
                        const newTutorials = subTutorials.filter(subTutorials => !prevTutorials.some(t => t.url === subTutorials.url));
                        return [...prevTutorials, ...newTutorials];
                    });
                    setPage(prevPage => prevPage + 1);
                } else {
                    setHasMore(false);
                }
            })
            .catch(err => {
                if (err.response) {
                    NotificationManager.error(`Data: ${err.response.data}\nStatus: ${err.response.status}\nHeaders: ${err.response.headers}`)
                } else if (err.request) {
                    NotificationManager.error(`${err.message}`)
                } else {
                    NotificationManager.error(`${err.message}`)
                }
                // NotificationManager.error("Something Went Wrong")
            });
    }

    const handleDelete = (tutorial) => {
        setTutorialToDelete(tutorial);
        setIsModalOpen(true);
    }
    const confirmDelete = () => {
        if (tutorialToDelete) {
            privateAxiosInstance.delete(`/tutorial/${tutorial_url}/${tutorialToDelete.url}`)
                .then(response => {
                    setTutorials(tutorials.filter(tutorial => tutorial.url !== tutorialToDelete.url));
                    setIsModalOpen(false);
                    setTutorialToDelete(null);
                })
                .catch(err => {
                    console.log("got error", err);
                });
        }
    }

    useEffect(() => {
        getTutorials();
    }, []);

    return (
        <div className="p-6">
            <h1 className="text-center text-2xl font-semibold">Manipulate Tutorial {tutorial_url}</h1>
            <div className='px-4 flex justify-end'>
                <Link to={`/kubari/admin/tutorials/edit/sub_tutorial/add/${tutorial_url}`} title='Add a sub tutorial'>
                    <PlusCircleIcon className="size-10 text-slate-600 hover:text-slate-700 cursor-pointer " />
                </Link>
            </div>
            <InfiniteScroll
                dataLength={tutorials.length}
                next={getTutorials}
                hasMore={hasMore}
                loader={<h4 className='text-center font-bold'>Loading...</h4>}
                endMessage={<p className="text-center font-bold">No more tutorials</p>}
            >
                <div className="tutorialSection p-3">
                    {tutorials.map(tutorial => (
                        <div key={tutorial.url} className="p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md md:flex md:justify-between">
                            <div className="contentDir">
                                <Link to={`/tutorials/${tutorial.url}`} className="text-sky-600 text-xl">
                                    <h1>{tutorial.title}</h1>
                                </Link>
                                <p className='text-sm mb-3'>{tutorial.author} . <span className='text-slate-400'>{formatDate(tutorial.created_at)} </span></p>
                                <p className='text-sm flex flex-wrap  space-y-2 items-center'>
                                    <span></span>
                                    {tutorial.tags.slice(0, 10).map(tag => (<p className='dark:bg-slate-700 bg-slate-300 p-1 mr-2  dark:text-slate-300 rounded-[10px] '>{tag}</p>))}
                                </p>
                            </div>
                            <div className="buttons flex">
                                <Link to={`/kubari/admin/tutorials/edit/sub_tutorial/${tutorial_url}/${tutorial.url}`}>
                                    <PencilSquareIcon className="size-9 text-slate-500 cursor-pointer" />
                                </Link>
                                <button onClick={() => handleDelete(tutorial)}>
                                    <TrashIcon className="size-9 text-slate-500 cursor-pointer mb-16" />
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </InfiniteScroll>
            <ConfirmationModal
                isOpen={isModalOpen}
                onClose={() => setIsModalOpen(false)}
                onConfirm={confirmDelete}
                message="Are you sure you want to delete this tutorial?"
            />
        </div>
    )
}

export default AdminGetAllSubTutorial