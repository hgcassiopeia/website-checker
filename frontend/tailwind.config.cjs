/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    fontFamily: {
      serif: ['Roboto']
    },
    extend: {
      colors: {
        'primary': '#6749F5',
        'gray-1': '#64697F',
        'gray-2': '#B6BED0',
        'gray-3': '#95A1BB',
        'gray-4': '#F5F7FC',
        'green-1': '#3FE364'
      }
    },
  },
  plugins: [],
};
