package templates

/* We assume that the table has unique non overlapping conditions. So maximal hit conditions can be one (or nothing).
   We cannot check the uniqueness. In case that the table is not unique, multi hits are solved with a priority hit policy
   (highest priority wins).

*/

const UNIQUE = ``
