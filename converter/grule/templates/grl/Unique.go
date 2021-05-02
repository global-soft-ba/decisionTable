package grl

/* We assume that the table has unique non overlapping conditions. So maximal hit conditions can be one (or nothing).
   We cannot check the uniqueness. In case that the table is not unique, multi hits are solved with a priority hit policy
   (highest priority wins).

*/

/* Would be correct for real unique => No Salience needed */
/* const UNIQUE = `{{define "SALIENCE"}} {{end}}` */

// UNIQUE Case with fallback in case of overlapping expressions
const UNIQUE = `{{define "SALIENCE"}}salience {{.Salience}} {{end}}`
