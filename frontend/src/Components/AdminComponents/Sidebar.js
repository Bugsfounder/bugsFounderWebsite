import React, { useState } from 'react';
import { Link } from 'react-router-dom';

function Sidebar({ sidebarOpen, toggleSidebar }) {
    return (
        <aside className={`${sidebarOpen ? 'w-60 sm:w-80' : 'w-16'} dark:text-slate-300 md:pt-7 fixed h-screen bg-gray-100 dark:bg-slate-900 flex flex-col px-5 shadow-lg dark:shadow-slate-500 shadow-slate-800 transition-all duration-300`}>
            <div className="sidebar">
                <div className="flex justify-between w-full">
                    <div className="hamburger cursor-pointer py-5" onClick={toggleSidebar}>
                        <div className="first dark:bg-slate-600 bg-black h-0.5 w-5 mb-1"></div>
                        <div className="second dark:bg-slate-600 bg-black h-0.5 w-5 mb-1"></div>
                        <div className="third dark:bg-slate-600 bg-black h-0.5 w-5"></div>
                    </div>
                    {sidebarOpen && (
                        <div className="titleSidebarlogo py-3">
                            <Link to="/">
                                <span className="font-extrabold text-xl" id="aboutLogo">
                                    <span className="dark:text-white text-gray-400">Bugs</span>
                                    <span className="text-sky-600">Founder</span>
                                </span>
                            </Link>
                        </div>
                    )}
                </div>
                {sidebarOpen && (
                    <div className="sidebarContent overflow-auto h-[500px]">
                        <div className="navItems mt-4 md:block md:mt-5">
                            <ul className="space-y-2 font-bold">
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/blogs">Blogs</Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/tutorials">Tutorials</Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/users">Users</Link></li>
                                <li><Link className="dark:hover:text-slate-400 md:text-lg" to="/kubari/admin/admins">Admins</Link></li>
                            </ul>
                        </div>
                    </div>
                )}
            </div>
        </aside>
    );
}

export default Sidebar;
