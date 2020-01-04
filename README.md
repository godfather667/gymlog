#GYMLOG

**NAME:**   gymlog - A new cli application

**USAGE:**   gymlog [global options] command [command options] [arguments...]

**COMMANDS**:

**gymlog page, p** ---Prints Page for Log Book: (Formats gymlog.ini for Log Book)

**gymlog data, d** ---Stores Page in Database: (Database Format)

**gymlog list, l** ---Lists Contents of Database by range:
list [ mm dd yyyy [ mm dd yyyy] } -No dates = all dates

**gymlog remove, r** ---Remove a record from Database:
Remove <record number> (see listing)

**gymloig chart, c** ---Produces Progress Chart:
chart [ mm dd yyyy [ mm dd yyyy ] ] -No dates = all dates

**help, h** ---Shows a list of commands or help for one command

GLOBAL OPTIONS:

**gymlog --help, -h**  show help (default: false)

**Initialization File Example:**

```
# INIT gymlog.ini         // (THIS FILE) Fixed Specification
# -------- FILE SPECIFICATIONS (Can be Changed)--------------
# PAGE pageFile.txt       // Listing Produced for Daily Workout
# DATA dataFile.txt       // Database File
# LIST listFile.txt       // Listing File - Holds produce listing
# ------------------ EXERCISE DATA --------------------------
#     |           |           |
#CODE |DESC       |EX-1       |EX-2
#=====|===========|===========|=========
MDY   Create_Date (12x23@2019)  
FWC   FW_Curl     (4x10@20)
FWT   FW_Tri      (4x10@30)
FWW   FW_Wrist    (4x15@20)
TRI   Tri_PullDwn (4x10@70)
LP    Leg_Press   (1x10@220)  (4x10@240)
LAT   Lat(L/R/B)  (4x10@90)
FLY   FLY(B/L/R)  (4x10@20)
LC    Leg_Curl    (1x10@55)   (4x10@70)
LE    Leg_Extend  (1x10@55)   (4x10@70)
ROW   Row         (4x10@85)
CP    Chest_Press (4x10@35)
CR    Crunches    (4x15@110)
LR    LAT_Raises  (4x10@20)
HT    HTub_Kick   (4x40)     
PW    Pool_Walk   (5x200)
PC    Pool_Calf   (4x20)        
```



