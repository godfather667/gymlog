﻿﻿﻿﻿﻿#GYMLOG**NAME:**   gymlog - A new cli application**USAGE:**   gymlog [global options] command [command options] [arguments...]**COMMANDS**:   **page, p**    Prints Page for Log Book: (Formats gymlog.ini for Log Book)   **data, d**    Store Page in Database: (Database Format)   **list, l**    List Contents of Database by range:              list mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed   **remove, r**  Remove a record from Database:             Remove <record number> (see FULL Listing For Record Number!!)  **chart, c**  Produces Progress Chart:              chart mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed   help, h  Shows a list of commands or help for one commandGLOBAL OPTIONS:**gymlog --help, -h**  show help (default: false)**Initialization File Example:**```# INIT gymlog.ini         // (THIS FILE) Fixed Specification# -------- FILE SPECIFICATIONS (Can be Changed)--------------# PAGE pageFile.txt       // Listing Produced for Daily Workout# DATA dataFile.txt       // Database File# LIST listFile.txt       // Listing File - Holds DataFile listing# CHART chartFile.txt     // Chart Listing - Holds Chart Listing# ------------------ EXERCISE DATA --------------------------#     |           |           |#CODE |DESC       |EX-1       |EX-2#=====|===========|===========|=========MDY   Create_Date (12x23@2019)  FWC   FW_Curl     (4x10@23)FWT   FW_Tri      (4x15@30)FWW   FW_Wrist    (4x15@23)TRI   Tri_PullDwn (4x10@40)TRP   Traps       (4x15@135)LP    Leg_Press   (1x10@220)  (4x10@240)LAT   Lat(L/R/B)  (4x10@90)FLY   FLY(B/L/R)  (4x10@20)LC    Leg_Curl    (1x10@70)   (4x10@85)LE    Leg_Extend  (1x10@70)   (4x10@85)ROW   Row         (4x10@85)CP    Chest_Press (4x10@35)CR    Crunches    (4x15@110)LR    LAT_Raises  (4x10@20)HT    HTub_Kick   (4x40)     PW    Pool_Walk   (5x200)PC    Pool_Calf   (4x20)- Comments preceded with a "#" Character- Three lines preceeded with "#" specify filenames for PAGE, DATA, LIST files- The filename for INIT is for documentation only -  It may not be changed```**USER GUIDE:**This program was a personal project to map my progress in weight training.  After 4/12 years in and out of a wheel chairdue to six knee operations, I had to rehab my body so I could return to a more normal life.The file *Initialization File (__gymlog.ini__)* contains:ged!**EXERCISE SECTION LAYOUT:**- Exercise Code- Exercise Description- Exercise 1 Specification- Exercise 2 Specification-**EXERCISE SPECIFICATION:**- "("- Cycle - How many times the reps and weights will be repeated- "x"- Reps - How many reps per weight- "@"- Weight - The weight to the nearest pound (or Kg)- ")"(4x10x70) Means  4 cycles of 10 times at 70lbs**One Note:** Some Exercises only have two entries; --- Example:  PW Pool Walk (5x200)  --- Some exercises have no weight: --- I walk  five laps of 200 yards in the shallow side of the Pool.In these cases the Reps are placed in the Weight Position in the chart.There are two Exercise Specifications.  These where primarily documenting any "Warm Ups" for a particular exercise. Only the exercise with the highest  number of reps will be used.**CHART COMMAND:**The chart lists exercises in database order.  ```Rec Date  |FWC |FWT |FWW |TRI |LP  |LAT |FLY |LC  |LE  |ROW |CP  |CR  |LR  |HT  |PW  |PC  |----------|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|1/5/2020  |20  |30  |20  |70  |240 |90  |20  |70  |70  |85  |35  |110 |20  |40  |200 |20  |----------|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|1/5/2020  |20  |30  |20  |70  |240 |90  |20  |70  |70  |85  |35  |110 |20  |40  |200 |20  |----------|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|1/5/2020  |20  |30  |20  |70  |240 |90  |20  |70  |70  |85  |35  |110 |20  |40  |200 |20  |----------|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|```Only the weights for the highest number of reps is shown! The columns identified by exercise codes.