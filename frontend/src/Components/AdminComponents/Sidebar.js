import React, { useState } from 'react';
import { Link } from 'react-router-dom';

function Sidebar() {
    const [sidebarOpen, setSidebarOpen] = useState(true);

    const toggleSidebar = () => {
        setSidebarOpen(!sidebarOpen);
    };

    return (
        <>
            <aside className={sidebarOpen ? ` dark:text-slate-300 md:pt-7 w-60 fixed h-screen bg-slate-200 dark:bg-slate-800 flex px-5 sm:w-80 shadow-lg  dark:shadow-slate-500 shadow-slate-800` : ` md:pt-7 `}>
                <div className="sidebar">
                    <div className={sidebarOpen ? `title flex justify-between w-52  sm:w-72` : `px-5`}>

                        <div className="hamburger cursor-pointer py-5" onClick={toggleSidebar}>
                            <div className="first dark:bg-slate-600 bg-black  h-0.5 w-5 mb-1"></div>
                            <div className="second dark:bg-slate-600 bg-black h-0.5 w-5 mb-1"></div>
                            <div className="third dark:bg-slate-600 bg-black  h-0.5 w-5 "></div>
                        </div>
                        <div className={sidebarOpen ? `titleSidebarlogo  py-3` : `hidden`}>
                            <Link to="/" className=''>
                                <span className='font-extrabold text-xl' id="aboutLogo"> <span className='dark:text-white text-gray-400'>Bugs</span>
                                    <span className='text-sky-600'>Founder</span>
                                </span>
                            </Link>
                        </div>


                    </div>

                    <div className={sidebarOpen ? `sidebarContent overflow-auto h-[500px]` : `hidden`}>

                        <div className="navItems mt-4 md:block md:mt-5 ">
                            <ul className='space-y-2 text-bold'>
                                <li></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/blogs">Blogs</Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/tutorials">Tutorials</Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/users">Users </Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/admins">Admins </Link></li>
                                <li></li>
                            </ul>
                        </div>
                    </div>
                </div>
            </aside>
        </>
    );
}

export default Sidebar;

// margin-left: 211px;