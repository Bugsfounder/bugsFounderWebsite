/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}",],
  darkMode: 'class',
  theme: {
    borderRadius: {
      'top-right-353': '353px',
    },
    extend: {
      width: {
        '2000': '2000px',
      },
    },
  },
  plugins: [],
}

