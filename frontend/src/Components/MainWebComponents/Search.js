import { useParams } from 'react-router'
import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
const Search = () => {
    const [loadedBlogs, setLoadedBlogs] = useState([]);
    const [page, setPage] = useState(1);
    const { query } = useParams();

    const allBlogs = [
        {
            id: "1",
            title: "The Future of Web Development",
            description: "Exploring the latest trends and technologies in web development for 2024.",
            author: "John Doe",
            createdAt: "2024-07-01",
            url: "/blogs/The-Future-of-Web-Development"
        },
        {
            id: "2",
            title: "Mastering React in 30 Days",
            description: "A comprehensive guide to becoming proficient in React.js within a month.",
            author: "Jane Smith",
            createdAt: "2024-06-15",
            url: "/blogs/Mastering-React-in-30-Days"
        },
        {
            id: "3",
            title: "Top 10 JavaScript Libraries",
            description: "An overview of the most popular JavaScript libraries and their use cases.",
            author: "Alex Johnson",
            createdAt: "2024-05-20",
            url: "/blogs/Top-10-JavaScript-Libraries"
        },
        {
            id: "4",
            title: "Understanding Async/Await",
            description: "A deep dive into asynchronous programming in JavaScript using async/await.",
            author: "Emily Davis",
            createdAt: "2024-04-10",
            url: "/blogs/Understanding-Async-Await"
        },
        {
            id: "5",
            title: "CSS Grid vs. Flexbox",
            description: "Comparing CSS Grid and Flexbox for modern web layout design.",
            author: "Michael Brown",
            createdAt: "2024-03-30",
            url: "/blogs/CSS-Grid-vs-Flexbox"
        },
        {
            id: "6",
            title: "Getting Started with TypeScript",
            description: "A beginner's guide to adding type safety to your JavaScript projects with TypeScript.",
            author: "Sarah Wilson",
            createdAt: "2024-02-25",
            url: "/blogs/Getting-Started-with-TypeScript"
        },
        {
            id: "7",
            title: "Building Progressive Web Apps",
            description: "Learn how to create fast, reliable, and engaging web apps using PWA principles.",
            author: "David Martinez",
            createdAt: "2024-01-15",
            url: "/blogs/Building-Progressive-Web-Apps"
        },
        {
            id: "8",
            title: "The Power of Serverless Architecture",
            description: "An introduction to serverless computing and its benefits for scalable applications.",
            author: "Sophia Garcia",
            createdAt: "2023-12-05",
            url: "/blogs/The-Power-of-Serverless-Architecture"
        },
        {
            id: "9",
            title: "GraphQL vs. REST API",
            description: "A comparison of GraphQL and REST for building robust APIs.",
            author: "James Taylor",
            createdAt: "2023-11-18",
            url: "/blogs/GraphQL-vs-REST-API"
        },
        {
            id: "10",
            title: "Optimizing Web Performance",
            description: "Techniques and best practices for improving the performance of your web applications.",
            author: "Olivia Moore",
            createdAt: "2023-10-22",
            url: "/blogs/Optimizing-Web-Performance"
        }
    ];

    const loadMoreBlogs = () => {
        // have to update this function when merging with backend to fetch data from backend
        const itemsPerPage = 3;
        const newBlogs = allBlogs.slice((page - 1) * itemsPerPage, page * itemsPerPage);
        setLoadedBlogs(prevBlogs => [...prevBlogs, ...newBlogs]);
    };

    useEffect(() => {
        loadMoreBlogs();
    }, [page]);

    useEffect(() => {
        const handleScroll = () => {
            if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
                setPage(prevPage => prevPage + 1);
            }
        };

        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, []);

    return (
        <div className="blogSection my-10 p-3 mt-[132.5px] container m-auto ">
            <div className='mx-5'>

                <h1 className="text-2xl font-bold my-8 dark:text-white">Search Results for: {query}</h1>
                {
                    loadedBlogs.map(blog => (
                        <div key={blog.id} className='p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md'>
                            <Link to={blog.url} className='text-sky-600 text-xl'>
                                <h1>{blog.title}</h1>
                            </Link>
                            <p>{blog.description}</p>
                            <p>{blog.author} . {blog.createdAt}</p>
                        </div>
                    ))
                }
                {loadedBlogs.length < allBlogs.length && <p className='text-gray-500 pl-1  hover:dark:text-gray-600'>Loading more blogs...</p>}
            </div>
        </div>
    );
}

export default Search