import React, { useEffect, useState } from 'react';
import { NotificationManager } from 'react-notifications';
import { useNavigate, useOutletContext, useParams } from 'react-router';
import { Link } from 'react-router-dom';
import DOMPurify from 'dompurify';
import hljs from 'highlight.js';
import 'highlight.js/styles/monokai.css';
import { calculateReadingTime } from '../utils/Utility';

const Tutorial = () => {
    const navigate = useNavigate();
    const { publicAxiosInstance } = useOutletContext();
    let { tutorial_url, sub_tutorial_url } = useParams();
    const [sidebarOpen, setSidebarOpen] = useState(window.innerWidth >= 768); // Initially open on larger screens
    const [tutorial, setTutorial] = useState({});
    const [subTutorial, setSubTutorial] = useState({});
    const [htmlContent, setHtmlContent] = useState("");
    const [readingTime, setReadingTime] = useState(0);
    const [activeSubTutorialUrl, setActiveSubTutorialUrl] = useState(sub_tutorial_url || "");

    const toggleSidebar = () => {
        setSidebarOpen(!sidebarOpen);
    };

    const getTutorial = (tutorial_url) => {
        publicAxiosInstance.get(`/tutorial/${tutorial_url}`)
            .then(response => {
                setTutorial(response.data);
                console.log("response.data: ", response.data);
                const initialSubTutorialUrl = sub_tutorial_url || response.data.sub_tutorials[0].url;
                setActiveSubTutorialUrl(initialSubTutorialUrl);
                getSubTutorial(tutorial_url, initialSubTutorialUrl);
            })
            .catch(err => {
                NotificationManager.error(err.message);
            });
    };

    const getSubTutorial = (tutorial_url, sub_tutorial_url) => {
        publicAxiosInstance.get(`/tutorial/${tutorial_url}/${sub_tutorial_url}`)
            .then(response => {
                setSubTutorial(response.data);
                const sanitizedHtml = DOMPurify.sanitize(response.data.content);
                setHtmlContent(sanitizedHtml);
                setReadingTime(calculateReadingTime(response.data.content.split(" ").length));
                navigate(`/tutorials/${tutorial_url}/${sub_tutorial_url}`, { replace: true });
            })
            .catch(err => {
                NotificationManager.error(err.message);
            });
    };

    useEffect(() => {
        getTutorial(tutorial_url);
    }, [tutorial_url]);

    useEffect(() => {
        if (sub_tutorial_url) {
            getSubTutorial(tutorial_url, sub_tutorial_url);
        }
    }, [sub_tutorial_url]);

    return (
        <div className="flex h-screen overflow-x-hidden">
            <aside className={`${sidebarOpen ? 'w-80' : 'w-16'} dark:text-slate-300 fixed h-full bg-gray-100 dark:bg-slate-900 flex flex-col px-5 shadow-lg dark:shadow-slate-500 shadow-slate-800 transition-all duration-300`}>
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
                        <div className="sidebarContent overflow-auto h-full">
                            <div className="navItems md:block md:mt-5 mt-4">
                                <ul className="space-y-2 font-bold">
                                    {tutorial.sub_tutorials && tutorial.sub_tutorials.map(sub_tutorial => (
                                        <li key={sub_tutorial.url}>
                                            <button
                                                className="dark:hover:text-slate-400 md:text-lg"
                                                onClick={() => getSubTutorial(tutorial_url, sub_tutorial.url)}
                                            >
                                                {sub_tutorial.title}
                                            </button>
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        </div>
                    )}
                </div>
            </aside>

            <div className={`flex-1 overflow-y-auto p-5 ${sidebarOpen ? 'ml-80' : 'ml-16'} transition-all duration-300`}>
                <div className="dark:text-gray-300 mt-10">
                    <div className="mb-10 dark:text-white">
                        <h1 className="text-center text-xl md:text-3xl lg:text-4xl font-semibold text-gray-800 dark:text-white mb-1 mt-[20px]">{subTutorial.title}</h1>
                        <div className="flex text-center items-center mb-6 flex-col md:flex-row md:justify-center">
                            {subTutorial.author} &nbsp; &middot;
                            <span className="text-slate-400">
                                &nbsp; {readingTime} min read
                            </span>
                        </div>
                        <div className="cont leading-relaxed text-dark dark:text-gray-100" dangerouslySetInnerHTML={{ __html: htmlContent }} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Tutorial;
