import LinkedIn, { GitHub, GMail, Telegram, Resume } from "../components/Icons";
// import {  Discord, Twitter } from "../components/Icons";
import ResumeLink from "../assets/resume.pdf"
const contactInfo = [
    {
        id: 1,
        name: "GitHub",
        link: "https://github.com/smilinh21",
        icon: <GitHub />,
    },
    {
        id: 2,
        name: "Mail",
        link: "mailto:lplinh0403@gmail.com",
        icon: <GMail />,
    },
    {
        id: 3,
        name: "Resume",
        link: ResumeLink,
        icon: <Resume />,
        download: true
    }
];

export default contactInfo;