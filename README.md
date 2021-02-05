# gogurls

URLs to get various GOG.com data.

## GOG.com URLs

Here are the key URLs functions exported by gogurls intended for client use:

### AccountProductsPage

- page
- media - gogtypes.Game, gogtypes.Movie
- sortBy - see gogtypes.AccountProductsSortOrder
- updated - return just the updated
- hidden - include items hidden by the user

### Details

- ID
- media - gogtypes.Game, gogtypes.Movie

### ProductsPage

- page
- media - gogtypes.Game, gogtypes.Movie
- sortOrder - see gogtypes.ProductsSortOrder

### WishlistPage

- page
- media - gogtypes.Game, gogtypes.Movie
- sortOrder - see gogtypes.WishlistSortOrder
- hidden - include items hidden by the user
