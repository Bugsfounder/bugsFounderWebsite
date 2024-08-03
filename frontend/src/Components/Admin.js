import React, { useState } from 'react';
import Navbar from './utils/Navbar';
import Sidebar from './AdminComponents/Sidebar';
import { Outlet } from 'react-router';

const Admin = () => {
    const [sidebarOpen, setSidebarOpen] = useState(true);

    const toggleSidebar = () => {
        setSidebarOpen(!sidebarOpen);
    };

    return (
        <>
            <Navbar />
            <div className="flex mt-[100.5px]">
                <Sidebar sidebarOpen={sidebarOpen} toggleSidebar={toggleSidebar} />
                <div className={`${sidebarOpen ? 'ml-60 sm:ml-80' : 'ml-16'} transition-all duration-300 w-full`}>
                    <div className="container m-auto">
                        <div className="dark:text-gray-300 mt-10">
                            <Outlet context={{}} />
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}

export default Admin;
