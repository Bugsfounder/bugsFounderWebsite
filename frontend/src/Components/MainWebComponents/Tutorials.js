import React, { useEffect, useState } from 'react';
import { Link, useOutletContext } from 'react-router-dom';
import InfiniteScroll from 'react-infinite-scroll-component';
import { NotificationManager } from 'react-notifications';


const Tutorials = () => {
    const { publicAxiosInstance } = useOutletContext();
    const [tutorials, setTutorials] = useState([]);
    const [page, setPage] = useState(1);
    const [hasMore, setHasMore] = useState(true);

    const getTutorials = () => {
        publicAxiosInstance.get(`/tutorial?page=${page}&limit=6`)
            .then(response => {
                console.log(response)
                if (response.data.tutorials && response.data.tutorials.length > 0) {
                    setTutorials(prevTutorials => {
                        const newTutorials = response.data.tutorials.filter(tutorial => !prevTutorials.some(b => b.url === tutorial.url));
                        console.log(response.data.tutorials)
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
            });
    }


    useEffect(() => {
        getTutorials();
    }, []);

    return (
        <div className="p-6  h-screen container m-auto ">
            <div className='mx-5'>

                <h1 className="text-2xl font-semibold">Tutorials</h1>
                <InfiniteScroll
                    dataLength={tutorials.length}
                    next={getTutorials}
                    hasMore={hasMore}
                    loader={<h4 className='text-center font-bold'>Loading...</h4>}
                    endMessage={<p className="text-center font-bold">No more tutorials</p>}
                >
                    <div className="tutorialSection py-3">
                        {tutorials.map(tutorial => (
                            <div key={tutorial.url} className="p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md md:flex md:justify-between" >
                                <div className="contentDir">
                                    <Link to={`/tutorials/${tutorial.url}/${tutorial.sub_tutorials[0].url}`} className="text-sky-600 text-xl">
                                        <h1>{tutorial.title}</h1>
                                    </Link>
                                    <p className='text-sm mb-3 text-slate-400'>{tutorial.author}
                                        {/* . <span className='text-slate-400'>{formatDate(tutorial.updated_at)} </span> */}
                                    </p>
                                    <p className='text-sm flex flex-wrap  space-y-2 items-center'>
                                        <span></span>
                                        {tutorial.tags.slice(0, 10).map(tag => (<p className='dark:bg-slate-700 bg-slate-300 p-1 mr-2  dark:text-slate-300 rounded-[10px] '>{tag}</p>))}
                                    </p>
                                </div>
                            </div>
                        ))}
                    </div>
                </InfiniteScroll>
            </div>
        </div>
    )
}

export default Tutorials