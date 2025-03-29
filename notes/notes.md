# Course Manager

## Developer Environment Dependencies

- Go (see go.mod for go dependencies)
- Docker
  - Marp runs in container
  - Caddy runs in container
- Sqlc CLI
- Task CLI
- Templ CLI
- Goose (for db migrations)
- Tailwind
- Typescript

## VS Code extensions used
- [Task](https://marketplace.visualstudio.com/items?itemName=task.vscode-task)
- [Task](https://marketplace.visualstudio.com/items?itemName=paulvarache.vscode-taskfile)
- [Tailwind CSS Intellisense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [Tailwind Docs](https://marketplace.visualstudio.com/items?itemName=austenc.tailwind-docs)
- [Tailwind Fold](https://marketplace.visualstudio.com/items?itemName=stivo.tailwind-fold)
- [Templ](https://marketplace.visualstudio.com/items?itemName=a-h.templ)

## Workflows

### Run in development (local)

Builds binary and mounts as volume, runs all services with docker compose:
- caddy
- marp
- echo app

```bash
task run-dev
```

To quickly regenerate templates and rebuild go code during development use:

```bash
task restart-echo
```

### Deploy

```bash
task deploy
```

### Run in production (remote, after deploy)

```bash
task run-prod
```

## Progress Log

### 3/29/25

- Had to write this down somewhere for now... currently my echo server is getting slides as a string from the marp server. But maybe I should instead just proxy my marp server with Caddy and put an iframe with the reverse proxy link to the slides? One big issue here is that the js in marp slides assumes the client is local host and it's trying to establish a websocket connection for live reload or something. So maybe I do need to filter that out with my app.

### 3/27/25

- Static site lesson details now uses tabs component to display slides, files, and assessments rather than trying to cram it all on one page. Looks great!
- Static site file viewer UI is looking great too, and the recursive part wasn't as bad as I expected!
- Realized I need to reorganize my project eventually to slice vertically according to features instead of horizontally according to layers. That seems to be the recommended approach nowadays compared with the older MVC approach. Of course this realization comes only after I have thoroughly entangled my features with each other according to layer (e.g. this monster NodeHandler interface)
- Pushed to remote in a matter of minutes and it works fantastic.

### 3/22/25

- Built a very smooth development and deployment setup this week with Docker compose including nice `task` shortcuts.
- Decided yesterday that I eventually want to make the following changes:
  - Use Kubernetes instead of Docker compose for production (probably overkill, but would facilitate scaling and need to learn it)
  - Postgres instead of sqlite3
  - Separate container for filesystem access to facilitate scaling up
  - ValKey (open source fork of Redis) for caching files
  - Nix instead of docker for development, possibly production as well
- But first I want to get this working completely for testing in beta so that I can share with a few select colleagues.
- Lots of changes since last entry. Today is a day to celebrate, as I finally got the whole thing working on my remote server. 3 containers running caddy, echo app, and marp respectively. Caddy is reverse proxy for echo and at {username}.course-manager.app caddy serves static sites.  It all works! I can't believe it! 🥳😭
- App is still registered as 'testing' not 'published' in google cloud console so I would need to add specific users until I publish it (authorized redirects cannot be http for published apps)
- Created unauthorized page and types and functions for responding to various unauthorized user scenarios
- Configured signin/signup route handlers to disallow non-cms gmail accounts

### 2/19/25

- Added the files capability to units and courses in the same way that it's available in lessons.
- I am increasingly noticing there's a tremendous potential for consolidating code to accommodate CourseNode interface and replacing the particular functions and data for course, unit, lesson, etc...
- This could cut down the total lines by at least 50%, if not more, and could make it more maintainable. I am very trepidatious about it though because it could be a lot of work and would have to be done with great care.

### 2/17/25

- Something interesting happened to me today. Nasty bug because I accidentally returned nil instead of err when err != nil, so it took me a while to find:

```json
{
  "error": "sql: Scan error on column index 5, name \"user_id\": converting driver.Value type float64 (\"1.0160211027267435e+20\") to a int64: invalid syntax",
  "message": "Internal Server Error"
}
```

- I am not sure what caused this, but hopefully will remember to write here what happened!
- I suspect it has something to do with my changing the id params type to an interface instead of integer but I thought any risk of doing so would be confined to the scope in which params was made; I did not think it could somehow infiltrate all the way down into SQL code.
- It looks like it's interpreting the user ID as a float. So that's an issue.  The bigger question is why is the userID getting passed in instead of the termID?
- Ok, the problem was very simple, and had nothing to do with the above. I hand't changed the Term's UserID field from int to string, so it was trying to read the number to an int. 🤦🏻‍♂️

### 2/16/25

- Deployed to Digital Ocean Droplet (see .env for IP)
- Purchased domain course-manager.app

### 2/6/2025

I've made a ton of additions/changes, just haven't been writing here lately.

I do want to mention that I spent several days exploring the possibility of using Hugo for site generation, but today I decided that having to use go templates vs. templ was not obviously  worth the benefits as far as I could see. Templ is just way, way easier to write UI with. The extra serialization required for using Hugo was also a major drag, having to pass data programmatically through markdown really did not seem worth it. A lot of my decision was just realizing the steepness of the learning curve involved. Templ is working just fine for me, and I just parallelized the site generation so that will help with the small performance concerns I had.

### 1/24/2025

Went a bit nuts today with the scale of changes.

Changed the filesystem to a nested one so that deleting a node (term, course, unit, lesson) will delete all of its children. For deleting things, I have a confirmation on the UI but maybe it's still a bit too easy to delete a lot of work, potentially forever!

Major change to static site in that I changed the directory names to match the server's new file structure so that I can simply copy the entire directory over once, including slides, images, and files, and then just render the pages. Vastly simpler than what I was doing before.

Removed images from DB, previously removed slides and files. So now a lot of things I had to register in the DB is simply housed within the node directory.

### 1/23/2025

Didn't accomplish as much today due to time constraints. Cleaned up NodeList and NodeDetails components code by adding corresponding methods to the PageData structs.

### 1/22/2025

Added ability to view and upload lesson files. Went a lot quicker than I expected.

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
