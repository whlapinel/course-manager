
# Action tree

```yml
Main Menu:

    View courses: lists all course templates without details
    
    Select course: lists terms for a given template
    
        View Details: lists all units and Lessons for a given course

            Edit Unit: turns all editable unit properties into form fields
        
            Edit Lesson: turns all editable lesson properties into form fields

            Quick Actions:

                Shift right: Change date of current lesson to subsequent day of instruction

                Shift left: Change date of current lesson to previous day of instruction
    
        Select Term: list course instance for the given course and term
```

## TODO

- 12/17/24 lesson struct should have Date field be Dates []time.Time instead of
Date time.Time so that a lesson can span multiple dates if necessary (not ideal
but sometimes necessary)
