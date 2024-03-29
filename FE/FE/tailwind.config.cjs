const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  important:true,
  theme: {
    extend: {},
    fontFamily: {
      signature: ["Arial"],
      advanced:['Nunito']
    },
    screens: {
      'xs': '475px',
      ...defaultTheme.screens,
    }
  },
  plugins: [],
}
