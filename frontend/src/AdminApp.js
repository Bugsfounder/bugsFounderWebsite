import React, { useState } from 'react';
import Navbar from './Components/utils/Navbar';
import Sidebar from './Components/utils/AdminSidebar';
import { Outlet } from 'react-router';
import axios from 'axios';

// Axios instance for public API
const publicAxiosInstance = axios.create({
    baseURL: 'http://localhost:8080/api/public',
    withCredentials: true, // If your backend requires credentials
});

// Axios instance for private API
const privateAxiosInstance = axios.create({
    baseURL: 'http://localhost:8080/api/private/admin',
    withCredentials: true, // If your backend requires credentials
});

const AdminApp = () => {
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
                    <div className="container m-auto h-screen">
                        <div className="dark:text-gray-300 mt-10">
                            <Outlet context={{
                                publicAxiosInstance,
                                privateAxiosInstance,
                            }} />
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}

export default AdminApp;
