from sqlalchemy import select, delete, update
from sqlalchemy.orm import sessionmaker


from schema import Project, Employee, EmployeeProject, Status, Position

def create_positions(base_session: sessionmaker, positions: list[Position]):
    with base_session() as session:
        for pos in positions:
            session.add(pos)

        session.commit()
    return True

def get_positions(base_session: sessionmaker):
    with base_session() as session:
        stmt = select(Position)
        res = list(session.execute(stmt).scalars().all())
        session.commit()
    return res

def create_statuses(base_session: sessionmaker, statuses: list[Status]):
    with base_session() as session:
        for status in statuses:
            session.add(status)

        session.commit()
    return True

def get_statuses(base_session: sessionmaker):
    with base_session() as session:
        stmt = select(Status)
        res = list(session.execute(stmt).scalars().all())
        session.commit()
    return res

def create_employee(assoc: EmployeeProject, base_session: sessionmaker, employee: Employee, project: Project):
    assoc.project = project
    employee.projects.append(assoc)

    with base_session() as session:
        session.add(employee)
        session.commit()
    return True

def get_employee(base_session: sessionmaker, id_: int) -> Employee:
    with base_session() as session:
        res = session.get_one(Employee, id_)
        session.commit()
    return res

def get_employees(base_session: sessionmaker) -> list[Employee]:
    with base_session() as session:
        stmt = select(Employee)
        res = list(session.execute(stmt).scalars().all())
        session.commit()
    return res

def delete_employee(base_session: sessionmaker, id_: int):
    with base_session() as session:
        stmt = delete(Employee).where(Employee.id == id_)
        session.execute(stmt)
        session.commit()
    return True

def get_project(base_session: sessionmaker,id_: int) -> Project:
    with base_session() as session:
        res = session.get_one(Project, id_)
        session.commit()
    return res

def get_projects(base_session: sessionmaker) -> list[Project]:
    with base_session() as session:
        stmt = select(Project)
        res = list(session.execute(stmt).scalars().all())
        session.commit()
    return res

def delete_project(base_session: sessionmaker, id_: int):
    with base_session() as session:
        stmt = delete(Project).where(Project.id == id_)
        session.execute(stmt)
        session.commit()
    return True

def delete_status(base_session: sessionmaker, id_: int):
    with base_session() as session:
        stmt = delete(Status).where(Status.id == id_)
        session.execute(stmt)
        session.commit()
    return True

def delete_position(base_session: sessionmaker, id_: int):
    with base_session() as session:
        stmt = delete(Position).where(Position.id == id_)
        session.execute(stmt)
        session.commit()
    return True


