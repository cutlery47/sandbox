from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from schema import Base

from tests import *

if __name__ == '__main__':
    engine = create_engine("sqlite://", echo=True)

    Base.metadata.create_all(engine)
    base_session = sessionmaker(bind=engine)

    test_create_statuses(base_session)
    test_get_statuses(base_session)

    test_delete_status(base_session)
    test_delete_position(base_session)

    test_create_positions(base_session)
    test_get_positions(base_session)

    test_create_employee(base_session)

    test_get_employees(base_session)
    test_get_employee(base_session)

    test_get_projects(base_session)
    test_get_project(base_session)

    test_delete_employee(base_session)

    "All tests have run successfully"

    Base.metadata.drop_all(engine)
