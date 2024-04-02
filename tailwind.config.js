/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    fontFamily: {
      Inter: ["Inter", "sans-serif"],
    },
    container: {
      center: true,
    },
    textColor: {
      primary: "#B6C4B6",
      secondary: "#EEF0E5",
    },
    extend: {
      colors: {
        primary: "#163020",
        secondary: "#304D30",
      },
    },
  },
  plugins: [],
};
