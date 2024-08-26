import React, { useState } from 'react';
import Logo from '../../images/banner.jpeg';
import { Link, useOutletContext } from 'react-router-dom';
import { SocialIcon } from 'react-social-icons';
import { NotificationManager } from 'react-notifications';

const Signup = () => {
  const { publicAxiosInstance } = useOutletContext();
  const [user, setUser] = useState({});
  const [isPresent, setIsPresent] = useState(false);

  const signupUser = (user) => {
    const config = {
      headers: {
        "Content-Type": "application/json"
      }
    };

    publicAxiosInstance.post('/user/signup', user, config)
      .then(response => {
        NotificationManager.success("Signup successful");
        console.log(response.data);
      })
      .catch(err => {
        NotificationManager.error(`${err.message}`);
      });
  };

  const isUsernameOrEmailUnique = (usernameOrEmail) => {
    const config = {
      headers: {
        "Content-Type": "application/json"
      }
    };

    const usernameOrEmailObj = { "usernameOrEmail": usernameOrEmail.toLowerCase() };
    publicAxiosInstance.post('/user/email_or_username_unique', usernameOrEmailObj, config)
      .then(response => {
        let isPresent = response.data.isPresent;
        setIsPresent(isPresent);
        if (isPresent) {
          NotificationManager.error("Username or email already exists");
        }
      })
      .catch(err => {
        NotificationManager.error(`${err.message}`);
      });
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setUser({
      ...user,
      [name]: value
    });
    if (name === "email" || name === "username") {
      isUsernameOrEmailUnique(value.toLowerCase());
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (!isPresent) {
      signupUser(user);
    }
  };

  return (
    <section className="bg-gray-50 dark:bg-gray-900 mt-[132.5px]">
      <div className="flex flex-col items-center justify-center md:px-4 md:py-20 py-8 mx-auto my-10">
        <Link to="/" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
          <img className="w-8 h-8 mr-2" src={Logo} alt="logo" />
          Bugs<span className='text-sky-600'>Founder</span>
        </Link>
        <div className="w-full bg-white rounded-[10px] shadow-md shadow-slate-300 dark:shadow-lg dark:shadow-slate-950 dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-slate-800">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
              Create an Account
            </h1>
            <form className="space-y-4 md:space-y-6" onSubmit={handleSubmit}>
              <div>
                <label htmlFor="username" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Unique Username</label>
                <input type="text" name="username" id="username" className="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" placeholder="Username123" required onChange={handleChange} />
              </div>
              <div>
                <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email Id</label>
                <input type="email" name="email" id="email" className="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" placeholder="name@company.com" required onChange={handleChange} />
              </div>
              <div>
                <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                <input type="password" name="password" id="password" placeholder="••••••••" className="bg-gray-50 border border-gray-300 text-gray-900 rounded-[10px] focus:ring-gray-600 focus:border-slate-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 outline-none" required onChange={handleChange} />
              </div>
              <button type="submit" className="w-full text-white bg-slate-500 hover:bg-slate-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-[10px] text-sm px-5 py-2.5 text-center dark:bg-slate-600 dark:hover:bg-slate-700 dark:focus:ring-slate-800">Sign Up</button>
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Already have an account? <Link to="/auth/login" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Login</Link>
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
  );
};

export default Signup;
