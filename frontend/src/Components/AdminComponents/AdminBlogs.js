import React, { useEffect, useState } from 'react';
import { Link, useOutletContext } from 'react-router-dom';
import { PencilSquareIcon, TrashIcon, PlusCircleIcon } from '@heroicons/react/24/solid';
import InfiniteScroll from 'react-infinite-scroll-component';
import ConfirmationModal from '../utils/ConfigmationModal';
import { NotificationManager } from 'react-notifications';
import formatDate from '../utils/Utility';

const AdminBlogs = () => {
  const { privateAxiosInstance, publicAxiosInstance } = useOutletContext();
  const [blogs, setBlogs] = useState([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [blogToDelete, setBlogToDelete] = useState(null);

  const getBlogs = () => {
    publicAxiosInstance.get(`/blog?page=${page}&limit=6`)
      .then(response => {
        if (response.data.blogs && response.data.blogs.length > 0) {
          setBlogs(prevBlogs => {
            const newBlogs = response.data.blogs.filter(blog => !prevBlogs.some(b => b.url === blog.url));
            return [...prevBlogs, ...newBlogs];
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

  const handleDelete = (blog) => {
    setBlogToDelete(blog);
    setIsModalOpen(true);
  }
  const confirmDelete = () => {
    if (blogToDelete) {
      privateAxiosInstance.delete(`/blog/${blogToDelete.url}`)
        .then(response => {
          setBlogs(blogs.filter(blog => blog.url !== blogToDelete.url));
          setIsModalOpen(false);
          setBlogToDelete(null);
        })
        .catch(err => {
          console.log("got error", err);
        });
    }
  }

  useEffect(() => {
    getBlogs();
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-center text-2xl font-semibold">Manipulate Blogs</h1>
      <div className='px-4 flex justify-end'>
        <Link to="add">
          <PlusCircleIcon className="size-10 text-slate-600 hover:text-slate-700 cursor-pointer " />
        </Link>
      </div>
      <InfiniteScroll
        dataLength={blogs.length}
        next={getBlogs}
        hasMore={hasMore}
        loader={<h4 className='text-center font-bold'>Loading...</h4>}
        endMessage={<p className="text-center font-bold">No more blogs</p>}
      >
        <div className="blogSection p-3">
          {blogs.map(blog => (
            <div key={blog.url} className="p-5 dark:text-white dark:bg-slate-800 bg-gray-100 my-3 shadow-sm rounded-md md:flex md:justify-between">
              <div className="contentDir">
                <Link to={`/blogs/${blog.url}`} className="text-sky-600 text-xl">
                  <h1>{blog.title}</h1>
                </Link>
                <p className='text-sm mb-3'>{blog.author} . <span className='text-slate-400'>{formatDate(blog.created_at)} </span></p>
                <p className='text-sm flex flex-wrap  space-y-2 items-center'>
                  <span></span>
                  {blog.tags.slice(0, 10).map(tag => (<p className='dark:bg-slate-700 bg-slate-300 p-1 mr-2  dark:text-slate-300 rounded-[10px] '>{tag}</p>))}
                </p>
              </div>
              <div className="buttons flex">
                <Link to={`edit/${blog.url}`}>
                  <PencilSquareIcon className="size-9 text-slate-500 cursor-pointer" />
                </Link>
                <button onClick={() => handleDelete(blog)}>
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
        message="Are you sure you want to delete this blog?"
      />
    </div>
  );
};

export default AdminBlogs;
