## deliverd
This is a project aimed to support my "Durham Snus" venture. It is a tracker and courier management system for packages dispatched. Basically, imagine the Royal Mail website. It is like this, but also in the backend includes the ability to:
- assign routes to couriers, and allow them to mark delivered
- add Drops through a secured API
- follow and update status of a package through multiple APIs (these being for different stages of the package's journey)
- etc

Presently open source, using a MySQL backend and Golang to program the API due to it's resilient and low-resource nature.

### Layout
The project is laid out in a way which makes sense to me, and not necessarily others, so see the layout:
```
server/
    backend/    - files for the module 'deliverd'
    ...         - these files are to run the module 'deliverd'
public/         - public facing content
```
Obviously, this is not complete and will be added to in later weeks/months.

### License
This code follows the "Creative Commons Attribution-NonCommerical-ShareAlike" (i.e. CC BY-NC-SA) license, basically outlining you may copy, redistribute, modify or utilise any of the code here, provided you credit, do not use it commerically and release any derivative work under the same license.

The code is copyright to Casn Corporation Limited, granted for usage to you given the above terms.