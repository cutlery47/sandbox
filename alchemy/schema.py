from sqlalchemy import ForeignKey, UniqueConstraint
from sqlalchemy.orm import Mapped, mapped_column, relationship, DeclarativeBase

from datetime import datetime

class Base(DeclarativeBase):
    pass

class Employee(Base):
    __tablename__ = 'employees'

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str]
    salary: Mapped[float]

    position: Mapped["Position"] = relationship(back_populates="employees")
    position_id: Mapped[int] = mapped_column(ForeignKey("positions.id"))

    projects: Mapped[list["EmployeeProject"]] = relationship(back_populates="employee")

class Project(Base):
    __tablename__ = 'projects'

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str]
    budget: Mapped[float]
    start_date: Mapped[datetime]
    end_date: Mapped[datetime]

    status: Mapped["Status"] = relationship(back_populates="projects")
    status_id: Mapped[int] = mapped_column(ForeignKey("statuses.id"))

    employees: Mapped[list["EmployeeProject"]] = relationship(back_populates="project")

class EmployeeProject(Base):
    __tablename__ = 'employee_projects'
    __table_args__ = (
        UniqueConstraint('employee_id', 'project_id', name='employee_project_uc'),
    )

    id: Mapped[int] = mapped_column(primary_key=True)

    employee: Mapped["Employee"] = relationship(back_populates="projects")
    employee_id: Mapped[int] = mapped_column(ForeignKey("employees.id"))

    project: Mapped["Project"] = relationship(back_populates="employees")
    project_id: Mapped[int] = mapped_column(ForeignKey("projects.id"))

class Position(Base):
    __tablename__ = 'positions'

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str]

    employees: Mapped["Employee"] = relationship(back_populates="position")

class Status(Base):
    __tablename__ = 'statuses'

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str]

    projects: Mapped["Project"] = relationship(back_populates="status")


