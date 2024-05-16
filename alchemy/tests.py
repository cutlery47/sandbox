from sqlalchemy.orm import sessionmaker
from sqlalchemy import select

from functions import create_positions, get_positions, delete_position
from functions import create_statuses, get_statuses, delete_status
from functions import create_employee, get_employee, get_employees, delete_employee
from functions import get_project, get_projects, delete_project

from datetime import datetime, timedelta

from schema import Position, Employee, Project, EmployeeProject, Status

positions = [Position(name="pos_1"),
             Position(name="pos_2"),
             Position(name="pos_3")]

statuses = [Status(name="stat_1"),
            Status(name="stat_2"),
            Status(name="stat_3")]

assoc = EmployeeProject()

def test_create_positions(base_session: sessionmaker):
    positions = [Position(name="pos_1"),
                 Position(name="pos_2"),
                 Position(name="pos_3")]
    res = create_positions(base_session, positions)
    assert res == True

def test_get_positions(base_session: sessionmaker):
    res = get_positions(base_session)
    assert len(res) == 3

def test_create_statuses(base_session: sessionmaker):
    statuses = [Status(name="stat_1"),
                 Status(name="stat_2"),
                 Status(name="stat_3")]
    res = create_statuses(base_session, statuses)
    assert res == True

def test_get_statuses(base_session: sessionmaker):
    res = get_statuses(base_session)
    assert len(res) == 3

def test_create_employee(base_session: sessionmaker):
    pos = positions[0]
    status = statuses[0]

    e = Employee(name="john",
                 salary=10,
                 position_id=pos.id,
                 position=pos)

    p = Project(name="project",
                budget=10,
                start_date=datetime.now(),
                end_date=datetime.now() + timedelta(days=1),
                status_id=status.id,
                status=status)

    res = create_employee(assoc=assoc, base_session=base_session, employee=e, project=p)
    assert res == True

def test_get_employees(base_session: sessionmaker):
    res = get_employees(base_session)
    assert len(res) != 0

def test_get_projects(base_session: sessionmaker):
    res = get_projects(base_session)
    assert len(res) != 0

def test_get_employee(base_session: sessionmaker):
    res = get_employee(base_session, 1)
    assert res is not None

def test_get_project(base_session: sessionmaker):
    res = get_project(base_session, 1)
    assert res is not None

def test_delete_status(base_session: sessionmaker):
    delete_status(base_session, 3)
    res = get_statuses(base_session)
    print(len(res))
    assert len(res) == 2

def test_delete_position(base_session: sessionmaker):
    delete_position(base_session, 3)
    res = get_positions(base_session)
    print(len(res))
    assert len(res) == 0

def test_delete_employee(base_session: sessionmaker):
    delete_employee(base_session, 1)
    res = get_employees(base_session)
    assert len(res) == 0
