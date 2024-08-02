import React from 'react'
import Logo from '../images/banner.jpeg'
import { Link } from 'react-router-dom'
import { SocialIcon } from 'react-social-icons'
const Signup = () => {
  return (
    <section class="bg-gray-50 dark:bg-gray-900 mt-[132.5px]">
      <div class="flex flex-col items-center justify-center md:px-4 md:py-20 py-8 mx-auto my-10">
        <Link to="/" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
          <img class="w-8 h-8 mr-2" src={Logo} alt="logo" />
          Bugs<span className='text-sky-600'>Founder</span>
        </Link>
        <div class="w-full bg-white rounded-[10px] shadow-md shadow-slate-300 dark:shadow-lg dark:shadow-slate-950 dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-slate-800">
          <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
              Create an Account
            </h1>
            <form class="space-y-4 md:space-y-6" action="#">
              <div>
                <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Unique Username</label>
                <input type="email" name="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" placeholder="Username123" required="" />
              </div>
              <div>
                <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email Id</label>
                <input type="email" name="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" placeholder="name@company.com" required="" />
              </div>
              <div>
                <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                <input type="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" required="" />
              </div>
              <button type="submit" class="w-full text-white bg-slate-500 hover:bg-slate-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-[10px] text-sm px-5 py-2.5 text-center dark:bg-slate-600 dark:hover:bg-slate-700 dark:focus:ring-slate-800">Sign Up</button>
              <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                Already have an account? <Link to="/auth/login" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Login</Link>
              </p>
            </form>
            <div className='dark: text-gray-500 bg-slate-500 h-[2px] flex items-center justify-center relative'>
              <span className='dark:bg-slate-800 bg-white px-2'>Or</span>
            </div>
            <div className="loginViaOthers flex flex-col justify-start space-y-2">
              <Link to="/">
                <div className='text-white bg-slate-500 rounded-[200px]'>
                  <SocialIcon url='google.com' bgColor='#334155' className='mr-3 ' />
                  Sign Up With Google
                </div>
              </Link>
              <Link to="/">
                <div className='text-white bg-slate-500 rounded-[200px]'>
                  <SocialIcon url='github.com' bgColor='#334155' className='mr-3 ' />
                  Sign Up With Github
                </div>
              </Link>

            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

export default Signup