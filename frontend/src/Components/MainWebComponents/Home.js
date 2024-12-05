import React, { useEffect, useState } from 'react';
import bannerImage from "../../images/banner.jpg";
import { Link, useOutletContext } from 'react-router-dom';
import { fetchBlogs, fetchTutorials } from '../utils/fetcher';

const Home = () => {
    const { publicAxiosInstance } = useOutletContext();
    const [blogs, setBlogs] = useState([]);
    const [tutorials, setTutorials] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const blogs = await fetchBlogs(publicAxiosInstance, 6);
                const topTutorials = await fetchTutorials(publicAxiosInstance, 6);
                setBlogs(blogs);
                setTutorials(topTutorials);
            } catch (err) {
                console.error("Error fetching data:", err);
            }
        };
        fetchData();
    }, [publicAxiosInstance]);
    return (
        <div className='container m-auto'>
            <div className="m-5">
                <div className="mt-[132.5px] welcome lg:flex order-0 pt2 bg-gray-100 dark:bg-slate-800 rounded-full">
                    <div>
                        <img src={bannerImage} alt="" className='w-2000 rounded-br-[500px]' />
                    </div>
                    <div className="right dark:text-white text-center lg:pr-10 p-5 flex flex-col justify-center items-center">
                        <h1 className='text-3xl lg:text-4xl font-bold'>Welcome To <span className='dark:text-white text-gray-400'>Bugs</span><span className='text-sky-600'>Founder</span></h1>
                        <p className='my-3'>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem molestiae rem corporis exercitationem? Minus obcaecati molestiae et esse, nesciunt provident neque laborum hic?</p>
                    </div>
                </div>

                <div className="blogSection">
                    <h1 className="text-2xl font-bold my-8 dark:text-white">Top Blogs</h1>
                    {blogs.length > 0 ? (
                        blogs.map(blog => (
                            <div key={blog.url} className='p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md'>
                                <Link to={`/blogs/${blog.url}`} className='text-sky-600 text-xl'>
                                    <h1>{blog.title}</h1>
                                </Link>
                                <p>{blog.description}</p>
                                <p>{blog.author} . {blog.createdAt}</p>
                            </div>
                        ))
                    ) : (
                        <p className="text-center dark:text-white">No Blogs</p>
                    )}
                    <Link to="/blogs/" className='text-gray-500 pl-1 hover:dark:text-gray-600'>More Blogs</Link>
                </div>

                <div className="tutorialSection">
                    <h1 className="text-2xl font-bold my-8 dark:text-white">Top Tutorials</h1>
                    {tutorials.length > 0 ? (
                        tutorials.map(tutorial => (
                            <div key={tutorial.url} className='p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md'>
                                <Link to={`/tutorials/${tutorial.url}`} className='text-sky-600 text-xl'>
                                    <h1>{tutorial.title}</h1>
                                </Link>
                                <p>{tutorial.description}</p>
                                <p>{tutorial.author} . {tutorial.createdAt}</p>
                            </div>
                        ))
                    ) : (
                        <p className="text-center dark:text-white">No Tutorials</p>
                    )}
                    <Link to="/blogs/" className='text-gray-500 pl-1 hover:dark:text-gray-600'>More Tutorials</Link>
                </div>
            </div>
        </div>
    );
};

export default Home;
