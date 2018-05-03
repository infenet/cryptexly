/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
function URLEntry(name, value) {
    Entry.call( this, ENTRY_TYPE_INPUT, name, value );
}

URLEntry.prototype = Object.create(Entry.prototype);
URLEntry.prototype.constructor = URLEntry;

URLEntry.prototype.Icon = function() {
    return 'globe';
}

URLEntry.prototype.Render = function(with_value){
    return this.formGroup( this.input('text', with_value) ); 
}
