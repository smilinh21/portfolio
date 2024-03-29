import React from 'react'
import Card from './Card';
import { Code,Link } from './Icons';
import CardCover from './CardCover';
const Project = ({ project: {name, img, code, description, style = { shadow: "shadow-gray-500", cover: "from-gray-500" } } }) => {
    return (
        <Card style={style} >

            <div className="group relative rounded-md cursor-pointer">
                <img src={`https://source.unsplash.com/random`} alt="Projec-Image" width="auto" height="auto" loading='lazy' title="Project" className='rounded-t-md w-60 h-60' />
                <CardCover text={description} style={style} />
            </div>
            <div className='flex justify-center items-center rounded-b-md'>
                <ProjectsLink name={name} style={"rounded-bl-md " + style.cover} link={code} />
            </div>
        </Card>
    )
}

export default Project;


const ProjectsLink = ({ name, link, style }) => {
    return (
        <a href={link} target="_blank" rel="noreferrer"
            className={` w-1/2 py-3 text-center  text-x flex justify-center items-center text-white`}>
            {name === "Code" ? <Code /> : <Link />} <span className='pl-2'>{name}</span>
        </a>
    )
}

