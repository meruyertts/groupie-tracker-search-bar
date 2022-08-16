# Description: #
The project consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

    It will be given an API(https://groupietrackers.herokuapp.com/api), that consists in four parts:

        The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

        The second one, locations, consists in their last and/or upcoming concert locations.

        The third one, dates, consists in their last and/or upcoming concert dates.

        And the last one, relation, does the link between all the other parts, artists, dates and locations.

Groupie tracker search bar consists of creating a functional program that searches, inside your website, for a specific text input. So the focus of this project is to create a way for the client to search a member or artist or any other attribute in the data system you made.

    The program should handle at least these search cases :
        artist/band name
        members
        locations
        first album date
        creation date
    The program must handle search input as case-insensitive.
    The search bar must have typing suggestions as you write.
        The search bar must identify and display in each suggestion the individual type of the search cases. (ex: Freddie Mercury -> member)
        For example if you start writing "phil" it should appear as suggestions Phil Collins - member and Phil Collins - artist/band. This is just an example of a display.

# Usage: # 
Run the program by typing go run cmd/main.go in the terminal
