import React from 'react'
import { Link } from 'react-router-dom'
import { SocialIcon } from 'react-social-icons'

const Footer = () => {
    return (
        <div className="footer  bg-slate-50 dark:bg-slate-800 p-10 ">
            <div className="socialMedia flex justify-between items-center">
                <div className="logo  flex justify-center items-center">
                    <Link to="/"><span className='font-extrabold text-xl' id="authorLogo"> <span className='dark:text-white text-gray-400'>Bugs</span><span className='text-sky-600'>Founder</span></span> </Link>
                    <span className='dark:text-white font-extralight text-4xl mx-3'>|</span>
                    <span className='text-slate-400'>Copyright © 2024 BugsFounder.com</span>
                </div>
                {/* <div className="social space-x-2">
                    <SocialIcon className='' bgColor="#334155" fgColor="white" url="https://www.instagram.com/" target="_blank" />
                    <SocialIcon bgColor="#334155" fgColor="white" url="https://www.linkedin.com/" target="_blank" />
                    <SocialIcon bgColor="#334155" fgColor="white" url="https://www.github.com/" target="_blank" />
                    <SocialIcon bgColor="#334155" fgColor="white" url="https://www.youtube.com/" target="_blank" />
                </div> */}
            </div>
        </div>
    )
}

export default Footer