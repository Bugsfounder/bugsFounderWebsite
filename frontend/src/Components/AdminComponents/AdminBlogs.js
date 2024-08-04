import React, { useEffect, useState, useCallback } from 'react';
import { Link, useOutletContext } from 'react-router-dom';
import { PencilSquareIcon, TrashIcon } from '@heroicons/react/24/solid';
import axios from 'axios';

// // Axios instance for public API
// const publicAxiosInstance = axios.create({
//   baseURL: 'http://localhost:8080/api/public',
//   withCredentials: true, // If your backend requires credentials
// });

// // Axios instance for private API
// const privateAxiosInstance = axios.create({
//   baseURL: 'http://localhost:8080/api/private/admin',
//   withCredentials: true, // If your backend requires credentials
// });

const AdminBlogs = () => {
  const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();

  const [loadedBlogs, setLoadedBlogs] = useState([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);

  const loadMoreBlogs = useCallback(async () => {
    try {
      console.log(`Fetching blogs for page: ${page}`);
      setHasMore(true); // If no more blogs are returned, set hasMore to false
      const response = await publicAxiosInstance.get(`/blog?page=${page}&limit=3`);
      const newBlogs = response.data.blogs;
      console.log(newBlogs);
      if (newBlogs.length > 0) {
        // Check for duplicates before updating state
        setLoadedBlogs(prevBlogs => {
          const blogIds = new Set(prevBlogs.map(blog => blog.id));
          const filteredNewBlogs = newBlogs.filter(blog => !blogIds.has(blog.id));
          return [...prevBlogs, ...filteredNewBlogs];
        });
        setPage(prevPage => prevPage + 1); // Increment page number only if new blogs are loaded
        setHasMore(false); // If no more blogs are returned, set hasMore to false
      } else {
        setHasMore(false); // If no more blogs are returned, set hasMore to false
      }
    } catch (error) {
      console.error('Error fetching blogs:', error);
    }
  }, [page]);

  useEffect(() => {
    loadMoreBlogs();
  }, []); // Call loadMoreBlogs only once on initial render

  useEffect(() => {
    const handleScroll = () => {
      if (window.innerHeight + window.scrollY >= document.body.offsetHeight && hasMore) {
        loadMoreBlogs(); // Load more blogs when scrolling to the bottom
      }
    };

    const debounceScroll = debounce(handleScroll, 200); // Debounce the scroll event handler

    window.addEventListener('scroll', debounceScroll);
    return () => window.removeEventListener('scroll', debounceScroll);
  }, [hasMore, loadMoreBlogs]);

  // Debounce function to limit the rate at which a function can fire
  function debounce(func, wait) {
    let timeout;
    return function (...args) {
      const later = () => {
        clearTimeout(timeout);
        func(...args);
      };
      clearTimeout(timeout);
      timeout = setTimeout(later, wait);
    };
  }

  return (
    <div className="p-6">
      <h1 className="text-center text-2xl font-semibold">Manipulate Blogs</h1>
      <div className="blogSection p-3">
        {loadedBlogs.map(blog => (
          <div key={blog.url} className="p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md md:flex md:justify-between">
            <div className="contentDir">
              <Link to={`/blogs/` + blog.url} className="text-sky-600 text-xl">
                <h1>{blog.title}</h1>
              </Link>
              <p>{blog.description}</p>
              <p>{blog.author} . {blog.createdAt}</p>
            </div>
            <div className="buttons flex">
              <Link to={`/kubari/admin/blogs/edit/${blog.url}`}>
                <PencilSquareIcon className="size-9 text-slate-500 cursor-pointer" />
              </Link>
              <Link to="/kubari/admin/blogs">
                <TrashIcon className="size-9 text-slate-500 cursor-pointer" />
              </Link>
            </div>
          </div>
        ))}
        {hasMore && <p className="text-gray-500 pl-1 hover:dark:text-gray-600 text-center">Loading more blogs...</p>}
      </div>
    </div>
  );
};

export default AdminBlogs;
