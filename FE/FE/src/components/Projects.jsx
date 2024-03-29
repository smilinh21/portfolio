import React from 'react'
import SectionHeading from './SectionHeading'
import Project from './Project'
import personalProjects from '../data/projects'
const Projects = () => {
    return (
        <div name="Projects" className='pt-10 bg-gradient-to-b from-black via-black to-gray-800 text-white'>
            <div className='section justify-between'>       
            <SectionHeading heading="Projects" secondHeading="Check out some of my work" />
                <div className='section flex items-center justify-center gap-10'>


                    {personalProjects.map((project) => {
                        return (
                            <Project key={project.name} project={project} />
                        )
                    })}

                </div>
            </div>
        </div>
    )
}

export default Projects
