import json
import os

def get_all_skills(content: dict) -> None:
    all_skills = set()
    
    for job in content.values():
        required_skills: dict = job['required_skills']
        for skill in required_skills.keys():
            all_skills.add(skill)  
    
    with open('./all_skills.txt', 'w') as w:
        w.write(':\n'.join(list(all_skills)))

def append_skills(content: dict) -> None:
    skills_dict: dict[str, list] = {}

    with open('./all_skills.txt', 'r') as file:
        skills = file.read().splitlines()
        for skill in skills:
            k, v = skill.split(':')
            if v == 'NONE': continue
            skills_dict[k] = v.split(',')
    
    for job, value in content.items():
        required_skills: dict = value['required_skills']
        for skill, courses in required_skills.items():
            relevant_course = skills_dict.get(skill)
            if relevant_course is None: continue
            if relevant_course is not None:
                if courses == 'None':
                    if len(relevant_course) == 1:
                        content[job]['required_skills'][skill] = relevant_course[0]
                    else: 
                        content[job]['required_skills'][skill] = relevant_course
                elif type(courses) is str:
                    content[job]['required_skills'][skill] = list(set([courses] + relevant_course))
                else:
                    new_courses = list(set(courses + relevant_course))
                    content[job]['required_skills'][skill] = new_courses
    
    with open('qualitative_analysis.json', 'w') as file:
        file.write(json.dumps(content, indent=4))
            

if __name__ == '__main__':
    with open('../Skill_Comparison/skills_by_course.json', 'r') as file:
        skills_file: dict = json.loads(file.read())
        
    # get_all_skills(skills_file)
    append_skills(skills_file)