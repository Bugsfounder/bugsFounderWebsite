import React from 'react'
import profileImage from '../../images/logo.jpeg'
import { SocialIcon } from 'react-social-icons'
const About = () => {
    return (
        <div className='container m-auto  '>
            <div className="m-5">

                <div className=" mt-[132.5px] about-page dark:bg-slate-900 dark:text-slate-300 p-6">
                    <h1 className="text-3xl font-bold mb-4">About Me</h1>

                    <div className="flex flex-col md:items-center mb-6">
                        <img src={profileImage} alt="Profile" className=" w-64 mb-4 rounded-[150px]" />
                        <p className="md:text-center">
                            Hi, I'm Manisha, a passionate programmer and software engineer with a love for coding and problem-solving.
                        </p>
                    </div>
                    {/* <h2 className="text-2xl font-semibold mb-2">Professional Background</h2>
                <p className="mb-4">
                    I hold a degree in [Your Degree] from [Your University]. Over the past [number of years], I've worked at [Company Names], where I have contributed to a variety of projects ranging from web development to database management.
                </p> */}
                    <h2 className="text-2xl font-semibold mb-2">Skills and Expertise</h2>
                    <ul className="list-disc list-inside mb-4">
                        <li>JavaScript (React, Next.js, Node.js)</li>
                        <li>Python (Django, Flask)</li>
                        <li>Go</li>
                        <li>HTML & CSS</li>
                        <li>C, C++</li>
                        <li>Database Management (SQL, MongoDB)</li>
                        <li>Version Control (Git)</li>
                    </ul>
                    <h2 className="text-2xl font-semibold mb-2">Projects and Achievements</h2>
                    <p className="mb-4">
                        I have worked on several exciting projects, including:
                    </p>
                    <ul className="list-disc list-inside mb-4">
                        <li>Developed a real-time chat application using React and Node.js.</li>
                        <li>Created a machine learning model for predicting housing prices using Python.</li>
                        <li>Built a responsive e-commerce website with a custom CMS.</li>
                        <li>Contributed to open-source projects on GitHub.</li>
                    </ul>
                    <h2 className="text-2xl font-semibold mb-2">Personal Interests</h2>
                    <p className="mb-4">
                        When I'm not coding, I enjoy:
                    </p>
                    <ul className="list-disc list-inside mb-4">
                        <li>Playing games, Like: BGMI, FreeFire, Chess and Ludo.</li>
                        <li>Exploring new technologies and staying updated with industry trends.</li>
                        <li>Reading tech blogs and books on software development.</li>
                        <li>Spending time with friends and family.</li>
                    </ul>
                    <h2 className="text-2xl font-semibold mb-2">Social Media</h2>
                    <p className="mb-4">
                        Feel free to get in touch with me at:
                        <p>Email: <a href="mailto:your-email@example.com" className="text-sky-600">bugsfounder2021@gmail.com</a></p>
                    </p>
                    <div className="social space-x-2">
                        <SocialIcon className='' bgColor="#334155" fgColor="white" url="https://www.instagram.com/bugsfounder" target="_blank" />
                        <SocialIcon bgColor="#334155" fgColor="white" url="https://www.linkedin.com/bugsfounder" target="_blank" />
                        <SocialIcon bgColor="#334155" fgColor="white" url="https://www.github.com/bugsfounder" target="_blank" />
                        <SocialIcon bgColor="#334155" fgColor="white" url="https://www.youtube.com/@bugsfounder" target="_blank" />
                    </div>
                </div>
            </div>
        </div>

    )
}

export default About