# public attributes vs private

class SimpleGradeBook(object):
    def __init__(self):
        self.grades = {}

    def add_student(self, name):
        self.grades[name] = []

    def report_grades(self, name, score):
        self.grades[name].append(score)

    def average_grade(self, name):
        grades = self.grades[name]
        return sum(grades) / len(grades)

book  = SimpleGradeBook()
book.add_student('Antony')
book.report_grades('Antony', 100)
book.report_grades('Antony', 40)

print(book.average_grade('Antony'))