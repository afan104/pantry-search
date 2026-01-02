### Frontend

- search: ingredient list, recipe url
- existing pantry

### Database

- schema:
  - ingredient
  - ingredientType (optional)
  - quantity
  - units
  - dateUpdated
  - expectedExpiry

### Functionality

- Can display entire pantry
- user can search pantry for specific ingredient(s) by specifying a list or a recipe url
- user can update/delete existing entries
  - when submitting a full recipe, can ask to deduct that specific amount from the pantry (and make additional adjustments before submitting)
- user can create new entries
- user can remove ingredient(s) in pantry
- website shows ingredients that are expiring soon/expired
  - expired ingredients that haven't been manually removed are removed from their sections and listed at the bottom in red (need to be thrown or update expiry)
  - expiring soon ingredients are listed at the top to show that they need to be eaten soon or thrown away
- uses standard/clear units and has conversion system in place. If input has no conversion to standard units, the standard units are displayed for manual comparison.
- tracks pantry state and makes backups for undoing

### APIs

- getIngredientsAll
- getIngredients
- postIngredient
- putIngredient
- deleteIngredient

### Future considerations

- user can input a file also
- add mobile compatibility
- track waste (sustainability)?
- integrate aws to get analytics info
