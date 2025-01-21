# Course Manager

## Progress Log

### 1/21/2025

Finally implemented Copying Course to Term. Thought I had trouble with units and lessons becoming orphaned after parent course was deleted but that doesn't seem to happen now. Something to watch out / test for though. It appears that all child rows are deleted by ON CASCADE, so that when a course is deleted, the units, lessons, and lesson_dates

### 1/18/2025

I've been making a ton of progress though I haven't been writing it here.

Major decision to be made regarding how to store term dates. Currently I have a table for instructional days and a table for non-instructional days but only one of them is used in the application, so the non-instruct table is empty. Since these days are mutually exclusive I have a challenge in that I need to make sure any change to one is checked against the other and handled appropriately.

So, say if user adds a non-instructional day for a date which currently is listed in the instructional days (a highly probable scenario), the date would need to be deleted along with any associated lesson_date rows. This would not delete the lesson since lesson_dates is just a junction table. But then dates.day_number would then need to be updated, and I'm not sure this makes sense. I think day_number should probably be removed from persistence and calculated on the fly.

So to recap, the plan is:
- Use both instruct_dates ("dates") and non_instruct dates and check instruct_dates whenever there is a change to non_instruct_dates.
- Remove day_number
  - Remove from domain struct
  - Apply migration
  - Modify schema to remove column from schema
  - Modify any application code that uses day_number (e.g. I think FitToTerm uses this)
  - I think that rather than fixing the import from csv code, I'll just delete it.
- Code to modify dates in either table should modify dates in the other table accordingly.
- Add ON DELETE CASCADE to lesson_dates if it's not already there
- Since day_number will be gone I'll need to ensure dates are sorted by date.

On second thought:
- I could pretty easily extrapolate the non-instruct days from the instruct days.
- Or I could put all the dates in one table, and add a bool column to indicate whether it's an instruct day. Could also add a reason column to indicate why it's a non-instruct day.

### 1/10/2025

Great progress on the web app. I've just about matched the feaures of the fyne app and about to surpass. Surprised at how easy it was to implement the Shift Lesson UI. Quality of everything is much better and I feel more confident in how everything is laid out. Having fun returning to web dev, and just absolutely love working with HTMX.

I happen to have made all my lesson names e.g. Lesson 1.2 but this is redundant info with the Lesson Number and Unit Number and makes the Name field/column pointless in my case.  A teacher would certainly to display the numbers, but they also might want a name for a lesson. I am thinking that perhaps I should display the lesson number but I should make the name separate, more like a name that briefly describes the topic or activity, less like the designation number. One problem is this would require fetching the unit number from the units table whenever I need to display a lesson, but that might not be such a pain if I just add it to the SQL queries.

### 1/9/2025

Decided to shift to building a web app instead of a Fyne app. One big reason is I want to be able to display slides in markdown and html side-by-side. It will also make deployment much more straightforward.

I am starting to think it was a mistake to decouple lessons, slides and files. Now I think I should just make copies of everything with each new quarter instead of trying to preserve space. So the slides and files will now be stored in lessons directory named by lessonID, and accessing outside the web app will still not be feasible but at least it will all be in the same place.

### 1/4/2025

Shifted to lesson planning today and had a lovely time working with the files through the app. It's still a little cumbersome.

### 1/3/2025

Big progress today. I'm now able to delete the courses directory and re-generate it from the app's database and file system. It's pretty fragile though, and there are some things that need to be fixed ASAP:
- Slides are tied directly with a lesson id, so the next semesters lessons won't inherit those slides. Big problem!
  - I'll need to figure out how to do this -- I think I will just do the same as what I've done with images and files, but I'm not sure if this is optimal or not.
- docs root folder is not generated, so js, styles etc. must be manipulated manually. only courses folder and below is generated. Need to make deleting this directory part of the generation code.
- Courses main page image is not an attribute of term, but should be (very minor issue)

### 1/2/2025

- Name-based directory naming
  - Pros: I can find a lesson outside the application
  - Cons: Potential for name collisions, caching builds is less feasible
- Id-based directory naming
  - Pros: Caching build is more feasible

### 1/1/2025

- If I change my site generator route naming to match the id-based structure, then the name-change possibility is not as much of a problem. Another problem with name-based structure is the potential for name conflicts - I need to protect against that somehow.
- But if I'm doing a full rebuild each time, the name-change possibility is no problem at all apart from the name-conflict issue. But the nested structure mitigates this risk since 2 lessons in the same unit having the same name is low-chance.

- I am flailing a bit today. Going back to work has me on the fence a little about what to work on. Sleep-deprived by myself, and nervous about not getting a job I've applied for.
- I'm nervous about messing something up right before going back to work so I put my most recent work on the file system in a new feature branch called file_feature. I'm nervous about making other improvements in the dev branch that might be hard to deconflict if and when the feature branch is complete.
- The questions that are plaguing me revolve around the same things mentioned on 12/30. I have begun implementing a File resource system but I have a lot of uncertainty around how this will work and integrate with the site generator. I am thinking that:
  - While the app will keep files using id-based filenames, the site generator will maintain the name-based structure, so when a lesson name is changed a new directory will be created and the files will be added. The old name will remain however which is ok if links no longer point to that name but what if there's eventually a name conflict because of these old files? Otherwise it is a problem from the perspective of cleanliness. How will I handle the orphan files? I can't build everything from scratch every time, it will take too long right?  Or maybe not, since the longest part is probably not the copying but the marp slide generation. If I run marp on the slides.md files in the file system instead of in the generator, maybe I could just copy the files every time I build the site. Still, we are talking about 700+ files minimum, for the lesson pages and slides alone, ignoring the other lesson files which are admittedly small in size and quantity. If I have to copy 800 to a thousand files every time I build the site, it could be a real drag.
  - Essentially what I need is a cache system, and I am really unsure about how to do that.
  - The files will reside in the file system and will be edited there through the app by the user, by checking out and checking back in or something like that. Not really sure how best to do it.
  - How will the markdown slides fit in?  Should they be a file resource or should it be written from the database?
  -

### 12/30/24

- 12/30/24 Today was a day of stepping back and realizing I have created a whole new mess of problems.
- 12/30/24 I am increasingly thinking that I should try to eliminate the central importance of the unit in the system.
  - A lesson should belong to a course rather than a unit
  - A course should have a slice of lessons rather than a slice of units
  - Parts of the course manager and course site that lists all units will also need to list lessons that don't have a unit
  - New queries to fetch lessons by unit id.
- 12/30/24 I realized today that allowing the user to change the name of a lesson would completely mess up the system, since the file system for the site generator depends on the names of lessons remaining constant across terms.
  - New hidden file system where directories match lessonIDs rather than lesson names.
  - User will no longer touch slide text directly but will rather edit this text indirectly either:
    1. Write to temporary file with file watcher and allow user to edit through another program
    2. Edit through the app, while (at some point) providing marp preview
  - Directory will be copied over for each course i.e. when a course is fit to a new term
  - Slide text will be stored in Database as text and added as a field of Lesson
  - Site generator will write the text to a markdown file before running marp to render to HTML
  - To save size requirements, images will be kept in a single directory so they don't need to be copied over for each course
  - If a lesson has a new ID
- 12/30/24 Worked on the edit lesson form a bit.
- 12/29/24 Major milestone today. Finally got the calendar interface working with the ability to move lessons one date left or right. Got rid of CourseSchedule across all layers.
