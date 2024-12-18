
# Action tree

```yml
Main Menu:

    View courses: lists all course templates without details (id, name, etc.)

        View Details: lists all units and Lessons for a given course
    
            Edit Course Unit: turns all editable unit properties into form fields

    Select course instance: lists terms for a given template
    
        Select Term: list course instance for the given course and term

        View Instance Details: lists all units and Lessons for a given course instance

            Edit Instance Unit: turns all editable unit properties into form fields
        
            Edit Instance Lesson: turns all editable lesson properties into form fields

            Quick Actions:

                Shift right: Change date of current lesson to subsequent day of instruction

                Shift left: Change date of current lesson to previous day of instruction
    
```

## PENDING

- 12/18/24 figure out how to relate templates to instances. should the template
id only be present for a row if they are marked by user as synchronized? i.e. 
until the user marks sync?

## COMPLETE

- 12/17/24 lesson struct should have Date field be Dates []time.Time instead of
Date time.Time so that a lesson can span multiple dates if necessary (not ideal
but sometimes necessary) (completed 12/17/24)