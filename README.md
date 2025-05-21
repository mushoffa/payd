# Payd Mini Project: Daily Worker Roster Management System

## Objective
Design and build a simple scheduling and roster management system for managing shifts of
daily workers. This simulates scheduling features used by companies with hourly/daily staff.

## Business Requirements
### Basic:
* Workers cannot request shifts already assigned to someone else
* No overlapping shift requests allowed per worker
* Max 1 shift per day, max 5 shifts per week per worker
* Admin can override or reassign approved shifts
* Conflict checking must occur on both worker request and admin approval
* Shift times are stored and compared in UTC

### Additional:
* Admin cannot request shifts past beyond today's date
* Minimum shift duration is set to 1 hour
* Maximum shift duration is set to 8 hour