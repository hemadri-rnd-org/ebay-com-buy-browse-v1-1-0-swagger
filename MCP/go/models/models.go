package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CompatibilityProperty represents the CompatibilityProperty schema from the OpenAPI specification
type CompatibilityProperty struct {
	Localizedname string `json:"localizedName,omitempty"` // The name of the product attribute that as been translated to the language of the site.
	Name string `json:"name,omitempty"` // The name of the product attribute, such as Make, Model, Year, etc.
	Value string `json:"value,omitempty"` // The value for the name attribute, such as BMW, R1200GS, 2011, etc.
}

// AttributeNameValue represents the AttributeNameValue schema from the OpenAPI specification
type AttributeNameValue struct {
	Name string `json:"name,omitempty"` // The name of the product attribute, such as Make, Model, Year, etc.
	Value string `json:"value,omitempty"` // The value for the name attribute, such as BMW, R1200GS, 2011, etc.
}

// TimeDuration represents the TimeDuration schema from the OpenAPI specification
type TimeDuration struct {
	Unit string `json:"unit,omitempty"` // An enumeration value that indicates the units (such as hours) of the time span. The enumeration value in this field defines the period of time being used to measure the duration. Valid Values: YEAR, MONTH, DAY, HOUR, CALENDAR_DAY, BUSINESS_DAY, MINUTE, SECOND, or MILLISECOND Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:TimeDurationUnitEnum'>eBay API documentation</a>
	Value int `json:"value,omitempty"` // Retrieves the duration of the time span (no units).The value in this field indicates the number of years, months, days, hours, or minutes in the defined period.
}

// ShippingOptionSummary represents the ShippingOptionSummary schema from the OpenAPI specification
type ShippingOptionSummary struct {
	Guaranteeddelivery bool `json:"guaranteedDelivery,omitempty"` // Indicates if the seller has committed to shipping the item with eBay Guaranteed Delivery. With eBay Guaranteed Delivery, the seller is committed to getting the line item to the buyer within 4 business days or less. See the Buying items with eBay Guaranteed Delivery help topic for more details about eBay Guaranteed Delivery.
	Maxestimateddeliverydate string `json:"maxEstimatedDeliveryDate,omitempty"` // The end date of the delivery window (latest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: For the best accuracy, always include the contextualLocation values in the X-EBAY-C-ENDUSERCTX request header.
	Minestimateddeliverydate string `json:"minEstimatedDeliveryDate,omitempty"` // The start date of the delivery window (earliest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: For the best accuracy, always include the contextualLocation values in the X-EBAY-C-ENDUSERCTX request header.
	Shippingcost ConvertedAmount `json:"shippingCost,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Shippingcosttype string `json:"shippingCostType,omitempty"` // Indicates the type of shipping used to ship the item. Possible values are FIXED (flat-rate shipping) and CALCULATED (shipping cost calculated based on item and buyer location).
}

// PaymentMethod represents the PaymentMethod schema from the OpenAPI specification
type PaymentMethod struct {
	Paymentinstructions []string `json:"paymentInstructions,omitempty"` // The payment instructions for the buyer, such as cash in person or contact seller.
	Paymentmethodbrands []PaymentMethodBrand `json:"paymentMethodBrands,omitempty"` // The payment method brands, including the payment method brand type and logo image.
	Paymentmethodtype string `json:"paymentMethodType,omitempty"` // The payment method type, such as credit card or cash. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PaymentMethodTypeEnum'>eBay API documentation</a>
	Sellerinstructions []string `json:"sellerInstructions,omitempty"` // The seller instructions to the buyer, such as accepts credit cards or see description.
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Regionname string `json:"regionName,omitempty"` // A localized text string that indicates the name of the region. Taxes are generally charged at the state/province level or at the country level in the case of VAT tax.
	Regiontype string `json:"regionType,omitempty"` // An enumeration value that indicates the type of region for the tax jurisdiction. Valid Values: STATE_OR_PROVINCE - Indicates the region is a state or province within a country, such as California or New York in the US, or Ontario or Alberta in Canada. COUNTRY - Indicates the region is a single country. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:RegionTypeEnum'>eBay API documentation</a>
}

// Taxes represents the Taxes schema from the OpenAPI specification
type Taxes struct {
	Shippingandhandlingtaxed bool `json:"shippingAndHandlingTaxed,omitempty"` // This indicates if tax is applied for the shipping cost.
	Taxjurisdiction TaxJurisdiction `json:"taxJurisdiction,omitempty"` // The type that defines the fields for the tax jurisdiction details.
	Taxpercentage string `json:"taxPercentage,omitempty"` // The percentage of tax.
	Taxtype string `json:"taxType,omitempty"` // This field indicates the type of tax that may be collected for the item. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:TaxType'>eBay API documentation</a>
	Ebaycollectandremittax bool `json:"ebayCollectAndRemitTax,omitempty"` // This field is only returned if true, and indicates that eBay will collect tax (sales tax, Goods and Services tax, or VAT) for at least one line item in the order, and remit the tax to the taxing authority of the buyer's residence.
	Includedinprice bool `json:"includedInPrice,omitempty"` // This indicates if tax was applied for the cost of the item.
}

// TaxJurisdiction represents the TaxJurisdiction schema from the OpenAPI specification
type TaxJurisdiction struct {
	Region Region `json:"region,omitempty"` // This type is used to provide region details for a tax jurisdiction.
	Taxjurisdictionid string `json:"taxJurisdictionId,omitempty"` // The identifier of the tax jurisdiction.
}

// UpdateCartItemInput represents the UpdateCartItemInput schema from the OpenAPI specification
type UpdateCartItemInput struct {
	Cartitemid string `json:"cartItemId,omitempty"` // The identifier of the item in the cart to be updated. This ID is generated when the item was added to the cart.
	Quantity int `json:"quantity,omitempty"` // The new quantity for the item that is being updated.
}

// MarketingPrice represents the MarketingPrice schema from the OpenAPI specification
type MarketingPrice struct {
	Discountamount ConvertedAmount `json:"discountAmount,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Discountpercentage string `json:"discountPercentage,omitempty"` // This field expresses the percentage of the seller discount based on the value in the originalPrice container.
	Originalprice ConvertedAmount `json:"originalPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Pricetreatment string `json:"priceTreatment,omitempty"` // Indicates the pricing treatment (discount) that was applied to the price of the item. Note: The pricing treatment affects the way and where the discounted price can be displayed. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceTreatmentEnum'>eBay API documentation</a>
}

// RatingHistogram represents the RatingHistogram schema from the OpenAPI specification
type RatingHistogram struct {
	Count int `json:"count,omitempty"` // The total number of user ratings that the product has received.
	Rating string `json:"rating,omitempty"` // This is the average rating for the product. As part of a product review, users rate the product. Products are rated from one star (terrible) to five stars (excellent), with each star having a corresponding point value - one star gets 1 point, two stars get 2 points, and so on. If a product had one four-star rating and one five-star rating, its average rating would be 4.5, and this is the value that would appear in this field.
}

// Seller represents the Seller schema from the OpenAPI specification
type Seller struct {
	Feedbackscore int `json:"feedbackScore,omitempty"` // The feedback score of the seller. This value is based on the ratings from eBay members that bought items from this seller.
	Selleraccounttype string `json:"sellerAccountType,omitempty"` // Indicates if the seller is a business or an individual. This is determined when the seller registers with eBay. If they register for a business account, this value will be BUSINESS. If they register for a private account, this value will be INDIVIDUAL. This designation is required by the tax laws in some countries. This field is returned only on the following sites. EBAY_AT, EBAY_BE, EBAY_CH, EBAY_DE, EBAY_ES, EBAY_FR, EBAY_GB, EBAY_IE, EBAY_IT, EBAY_PL Valid Values: BUSINESS or INDIVIDUAL Code so that your app gracefully handles any future changes to this list.
	Username string `json:"username,omitempty"` // The user name created by the seller for use on eBay.
	Feedbackpercentage string `json:"feedbackPercentage,omitempty"` // The percentage of the total positive feedback.
}

// CompatibilityResponse represents the CompatibilityResponse schema from the OpenAPI specification
type CompatibilityResponse struct {
	Compatibilitystatus string `json:"compatibilityStatus,omitempty"` // An enumeration value that tells you if the item is compatible with the product. The values are: COMPATIBLE - Indicates the item is compatible with the product specified in the request. NOT_COMPATIBLE - Indicates the item is not compatible with the product specified in the request. Be sure to check all the value fields to ensure they are correct as errors in the value can also cause this response. UNDETERMINED - Indicates one or more attributes for the specified product are missing so compatibility cannot be determined. The response returns the attributes that are missing. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityStatus'>eBay API documentation</a>
	Warnings []Error `json:"warnings,omitempty"` // An array of warning messages. These types of errors do not prevent the method from executing but should be checked.
}

// AdditionalProductIdentity represents the AdditionalProductIdentity schema from the OpenAPI specification
type AdditionalProductIdentity struct {
	Productidentity []ProductIdentity `json:"productIdentity,omitempty"` // An array of the product identifier/value pairs for the product associated with the item. This is returned if the seller has associated the eBay Product Identifier (ePID) with the item and the request has fieldgroups set to PRODUCT. The following table shows what is returned, based on the item information provided by the seller, when the fieldgroups set to PRODUCT. ePID Provided Product&nbsp;ID(s) Provided Response No No The AdditionalProductIdentity container is not returned. No Yes The AdditionalProductIdentity container is not returned but the product identifiers specified by the seller are returned in the localizedAspects container. Yes No The AdditionalProductIdentity container is returned listing the product identifiers of the product. Yes Yes The AdditionalProductIdentity container is returned listing all the product identifiers of the product and the product identifiers specified by the seller are returned in the localizedAspects container.
}

// RemoveCartItemInput represents the RemoveCartItemInput schema from the OpenAPI specification
type RemoveCartItemInput struct {
	Cartitemid string `json:"cartItemId,omitempty"` // The identifier of the item in the cart to be removed. This ID is generated when the item was added to the cart.
}

// Price represents the Price schema from the OpenAPI specification
type Price struct {
	Convertedfromvalue string `json:"convertedFromValue,omitempty"` // The monetary amount before any conversion is performed, in the currency specified by the convertedFromCurrency field. This value is the pre-conversion amount. The value field contains the converted amount of this value, in the currency specified by the currency field.
	Currency string `json:"currency,omitempty"` // The three-letter ISO 4217 code representing the currency of the amount in the value field. If currency conversion/localization was performed, this is the post-conversion currency of the amount in the value field. Default: The currency of the user's country. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
	Value string `json:"value,omitempty"` // The amount of the currency specified in the currency field. The value of currency defaults to the standard currency used by the country of the eBay site offering the item. If currency conversion/localization was performed, this is the post-conversion amount. Default: The currency of the user's country.
	Convertedfromcurrency string `json:"convertedFromCurrency,omitempty"` // The three-letter ISO 4217 code representing the currency of the amount in the convertedFromValue field. This value is the pre-conversion currency. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
}

// ReviewRating represents the ReviewRating schema from the OpenAPI specification
type ReviewRating struct {
	Averagerating string `json:"averageRating,omitempty"` // The average rating given to a product based on customer reviews.
	Ratinghistograms []RatingHistogram `json:"ratingHistograms,omitempty"` // An array of containers for the product rating histograms that shows the review counts and the product rating.
	Reviewcount int `json:"reviewCount,omitempty"` // The total number of reviews for the item.
}

// ProductIdentity represents the ProductIdentity schema from the OpenAPI specification
type ProductIdentity struct {
	Identifiervalue string `json:"identifierValue,omitempty"` // The product identifier value.
	Identifiertype string `json:"identifierType,omitempty"` // The type of product identifier, such as UPC and EAN.
}

// ItemGroup represents the ItemGroup schema from the OpenAPI specification
type ItemGroup struct {
	Commondescriptions []CommonDescriptions `json:"commonDescriptions,omitempty"` // An array of containers for a description and the item IDs of all the items that have this exact description. Often the item variations within an item group all have the same description. Instead of repeating this description in the item details of each item, a description that is shared by at least one other item is returned in this container. If the description is unique, it is returned in the items.description field.
	Items []Item `json:"items,omitempty"` // An array of containers for all the item variation details, excluding the description.
	Warnings []Error `json:"warnings,omitempty"` // An array of warning messages. These types of errors do not prevent the method from executing but should be checked.
}

// ConvertedAmount represents the ConvertedAmount schema from the OpenAPI specification
type ConvertedAmount struct {
	Convertedfromvalue string `json:"convertedFromValue,omitempty"` // The monetary amount before any conversion is performed, in the currency specified by the convertedFromCurrency field. This value is required or returned only if currency conversion/localization is required. The value field contains the converted amount of this value, in the currency specified by the currency field.
	Currency string `json:"currency,omitempty"` // The three-letter ISO 4217 code representing the currency of the amount in the value field. If currency conversion/localization is required, this is the post-conversion currency of the amount in the value field. Default: The currency of the authenticated user's country. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
	Value string `json:"value,omitempty"` // The monetary amount, in the currency specified by the currency field. If currency conversion/localization is required, this value is the converted amount, and the convertedFromValue field contains the amount in the original currency.
	Convertedfromcurrency string `json:"convertedFromCurrency,omitempty"` // The three-letter ISO 4217 code representing the currency of the amount in the convertedFromValue field. This value is required or returned only if currency conversion/localization is required, and represents the pre-conversion currency. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Categoryid string `json:"categoryId,omitempty"` // The unique identifier of the primary item category of the item, as well as the secondary item category if item was listed in two categories.
}

// SearchPagedCollection represents the SearchPagedCollection schema from the OpenAPI specification
type SearchPagedCollection struct {
	Autocorrections AutoCorrections `json:"autoCorrections,omitempty"`
	Limit int `json:"limit,omitempty"` // The value of the limit parameter submitted in the request, which is the maximum number of items to return on a page, from the result set. A result set is the complete set of items returned by the method.
	Offset int `json:"offset,omitempty"` // This value indicates the offset used for current page of items being returned. Assume the initial request used an offset of 0 and a limit of 3. Then in the first page of results, this value would be 0, and items 1-3 are returned. For the second page, this value is 3 and so on.
	Prev string `json:"prev,omitempty"` // The URI for the previous page of results. This is returned if there is a previous page of results from the result set. The following example of the search method returns items 1 thru 5 from the list of items found, which would be the first set of items returned. https://api.ebay.com/buy/v1/item_summary/search?query=t-shirts&amp;limit=5&amp;offset=0
	Href string `json:"href,omitempty"` // The URI of the current page of results. The following example of the search method returns items 1 thru 5 from the list of items found. https://api.ebay.com/buy/v1/item_summary/search?q=shirt&amp;limit=5&amp;offset=0.
	Refinement Refinement `json:"refinement,omitempty"` // This type defines the fields for the various refinements of an item. You can use the information in this container to create histograms, which help shoppers choose exactly what they want.
	Next string `json:"next,omitempty"` // The URI for the next page of results. This value is returned if there is an additional page of results to return from the result set. The following example of the search method returns items 5 thru 10 from the list of items found. https://api.ebay.com/buy/v1/item_summary/search?query=t-shirts&amp;limit=5&amp;offset=10
	Total int `json:"total,omitempty"` // The total number of items that match the input criteria.
	Warnings []Error `json:"warnings,omitempty"` // The container with all the warnings for the request.
	Itemsummaries []ItemSummary `json:"itemSummaries,omitempty"` // An array of the items on this page. The items are sorted according to the sorting method specified in the request.
}

// VatDetail represents the VatDetail schema from the OpenAPI specification
type VatDetail struct {
	Vatid string `json:"vatId,omitempty"` // The seller's VAT (value added tax) ID. VAT is a tax added by some European countries.
	Issuingcountry string `json:"issuingCountry,omitempty"` // The two-letter ISO 3166 standard of the country issuing the seller's VAT (value added tax) ID. VAT is a tax added by some European countries. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	Image Image `json:"image,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
	Mpns []string `json:"mpns,omitempty"` // An array of all possible MPN values associated with the product. A MPNs is manufacturer part number of the product. To identify the product, this is always used along with brand.
	Aspectgroups []AspectGroup `json:"aspectGroups,omitempty"` // An array of containers for the product aspects. Each group contains the aspect group name and the aspect name/value pairs.
	Brand string `json:"brand,omitempty"` // The brand associated with product. To identify the product, this is always used along with MPN (manufacturer part number).
	Description string `json:"description,omitempty"` // The rich description of an eBay product, which might contain HTML.
	Gtins []string `json:"gtins,omitempty"` // An array of all the possible GTINs values associated with the product. A GTIN is a unique Global Trade Item number of the item as defined by https://www.gtin.info. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value.
	Additionalimages []Image `json:"additionalImages,omitempty"` // An array of containers with the URLs for the product images that are in addition to the primary image.
	Title string `json:"title,omitempty"` // The title of the product.
	Additionalproductidentities []AdditionalProductIdentity `json:"additionalProductIdentities,omitempty"` // An array of product identifiers associated with the item. This container is returned if the seller has associated the eBay Product Identifier (ePID) with the item and in the request fieldgroups is set to PRODUCT.
}

// AddCartItemInput represents the AddCartItemInput schema from the OpenAPI specification
type AddCartItemInput struct {
	Itemid string `json:"itemId,omitempty"` // The eBay RESTful identifier of the item you want added to the cart. RESTful Item ID Format: v1|#|# For example: v1|272394640372|0 v1|162846450672|461882996982 For more information about item ID for RESTful APIs, see the Legacy API compatibility section of the Buy APIs Overview. Maximum number of items in a cart: 100
	Quantity int `json:"quantity,omitempty"` // The number of this item the buyer wants to purchase. If this value is greater than the number available, the service will change this value to the number available. If this happens, a warning is returned. Maximum: number available
}

// Item represents the Item schema from the OpenAPI specification
type Item struct {
	Paymentmethods []PaymentMethod `json:"paymentMethods,omitempty"` // The payment methods for the item, including the payment method types, brands, and instructions for the buyer.
	Agegroup string `json:"ageGroup,omitempty"` // (Primary Item Aspect) The age group for which the product is recommended. For example, newborn, infant, toddler, kids, adult, etc. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Itemid string `json:"itemId,omitempty"` // The unique RESTful identifier of the item.
	Energyefficiencyclass string `json:"energyEfficiencyClass,omitempty"` // This indicates the European energy efficiency rating (EEK) of the item. This field is returned only if the seller specified the energy efficiency rating. The rating is a set of energy efficiency classes from A to G, where 'A' is the most energy efficient and 'G' is the least efficient. This rating helps buyers choose between various models. When the manufacturer's specifications for this item are available, the link to this information is returned in the productFicheWebUrl field.
	Brand string `json:"brand,omitempty"` // (Primary Item Aspect) The name brand of the item, such as Nike, Apple, etc. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Adultonly bool `json:"adultOnly,omitempty"` // This indicates if the item is for adults only. For more information about adult-only items on eBay, see Adult items policy for sellers and Adult-Only items on eBay for buyers.
	Primaryitemgroup ItemGroupSummary `json:"primaryItemGroup,omitempty"` // The type that defines the fields for the details of each item in an item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. When an item group is created, one of the item variations, such as the red shirt size L, is chosen as the "parent". All the other items in the group are the children, such as the blue shirt size L, red shirt size M, etc. <br /><br /><span class="tablenote"><b> Note: </b> This container is returned only if the <b> item_id</b> in the request is an item group (parent ID of an item with variations).</span>
	Seller SellerDetail `json:"seller,omitempty"` // The type that defines the fields for basic and detailed information about the seller of the item returned by the <b> item</b> resource.
	Warnings []Error `json:"warnings,omitempty"` // An array of warning messages. These types of errors do not prevent the method from executing but should be checked.
	Pattern string `json:"pattern,omitempty"` // (Primary Item Aspect) Text describing the pattern used on the item. For example, paisley. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Gender string `json:"gender,omitempty"` // (Primary Item Aspect) The gender for the item. This is used for items that could vary by gender, such as clothing. For example: male, female, or unisex. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Buyingoptions []string `json:"buyingOptions,omitempty"` // A comma separated list of all the purchase options available for the item. The values returned are: FIXED_PRICE - Indicates the buyer can purchase the item for a set price using the Buy It Now button. AUCTION - Indicates the buyer can place a bid for the item. After the first bid is placed, this becomes a live auction item and is the only buying option for this item. BEST_OFFER - Indicates the buyer can send the seller a price they're willing to pay for the item. The seller can accept, reject, or send a counter offer. For more information on how this works, see Making a Best Offer. Code so that your app gracefully handles any future changes to this list.
	Size string `json:"size,omitempty"` // (Primary Item Aspect) The size of the item. For example, '7' for a size 7 shoe. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Bidcount int `json:"bidCount,omitempty"` // This integer value indicates the total number of bids that have been placed against an auction item. This field is returned only for auction items.
	Categorypath string `json:"categoryPath,omitempty"` // Text that shows the category hierarchy of the item. For example: Computers/Tablets &amp; Networking, Laptops &amp; Netbooks, PC Laptops &amp; Netbooks
	Quantitylimitperbuyer int `json:"quantityLimitPerBuyer,omitempty"` // The maximum number for a specific item that one buyer can purchase.
	Availablecoupons []AvailableCoupon `json:"availableCoupons,omitempty"` // A list of available coupons for the item.
	Itemenddate string `json:"itemEndDate,omitempty"` // The date and time up to which the items can be purchased. This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: This field is not returned for Good 'Til Cancelled (GTC) listings.
	Itemlocation Address `json:"itemLocation,omitempty"` // The type that defines the fields for an address.
	Returnterms ItemReturnTerms `json:"returnTerms,omitempty"` // The type that defines the fields for the seller's return policy.
	Description string `json:"description,omitempty"` // The full description of the item that was created by the seller. This can be plain text or rich content and can be very large.
	Shiptolocations ShipToLocations `json:"shipToLocations,omitempty"` // The type that defines the fields that include and exclude geographic regions affecting where the item can be shipped. The seller defines these regions when listing the item.
	Localizedaspects []TypedNameValue `json:"localizedAspects,omitempty"` // An array of containers that show the complete list of the aspect name/value pairs that describe the variation of the item.
	Estimatedavailabilities []EstimatedAvailability `json:"estimatedAvailabilities,omitempty"` // The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned.
	Marketingprice MarketingPrice `json:"marketingPrice,omitempty"` // The type that defines the fields that describe a seller discount.
	Color string `json:"color,omitempty"` // (Primary Item Aspect) Text describing the color of the item. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Unitpricingmeasure string `json:"unitPricingMeasure,omitempty"` // The designation, such as size, weight, volume, count, etc., that was used to specify the quantity of the item. This helps buyers compare prices. For example, the following tells the buyer that the item is 7.99 per 100 grams. &quot;unitPricingMeasure&quot;: &quot;100g&quot;, &quot;unitPrice&quot;: { &nbsp;&nbsp;&quot;value&quot;: &quot;7.99&quot;, &nbsp;&nbsp;&quot;currency&quot;: &quot;GBP&quot;
	Conditionid string `json:"conditionId,omitempty"` // The identifier of the condition of the item. For example, 1000 is the identifier for NEW. For a list of condition names and IDs, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Shortdescription string `json:"shortDescription,omitempty"` // This text string is derived from the item condition and the item aspects (such as size, color, capacity, model, brand, etc.).
	Reservepricemet bool `json:"reservePriceMet,omitempty"` // This indicates if the reserve price of the item has been met. A reserve price is set by the seller and is the minimum amount the seller is willing to sell the item for. If the highest bid is not equal to or higher than the reserve price when the auction ends, the listing ends and the item is not sold. Note: This is returned only for auctions that have a reserve price.
	Shippingoptions []ShippingOption `json:"shippingOptions,omitempty"` // An array of shipping options containers that have the details about cost, carrier, etc. of one shipping option.
	Itemweburl string `json:"itemWebUrl,omitempty"` // The URL of the View Item page of the item. This enables you to include a &quot;Report Item on eBay&quot; link that takes the buyer to the View Item page on eBay. From there they can report any issues regarding this item to eBay.
	Minimumpricetobid ConvertedAmount `json:"minimumPriceToBid,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Sizetype string `json:"sizeType,omitempty"` // (Primary Item Aspect) Text describing a size group in which the item would be included, such as regular, petite, plus, big-and-tall or maternity. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Eligibleforinlinecheckout bool `json:"eligibleForInlineCheckout,omitempty"` // This field indicates if the item can be purchased using the Buy Order API. If the value of this field is true, this indicates that the item can be purchased using the Order API. If the value of this field is false, this indicates that the item cannot be purchased using the Order API and must be purchased on the eBay site.
	Epid string `json:"epid,omitempty"` // An EPID is the eBay product identifier of a product from the eBay product catalog. This indicates the product in which the item belongs.
	Product Product `json:"product,omitempty"` // The type that defines the fields for the product information of the item.
	Sizesystem string `json:"sizeSystem,omitempty"` // (Primary Item Aspect) The sizing system of the country. All the item aspects, including this aspect, are returned in the localizedAspects container. Valid Values: AU (Australia), BR (Brazil), CN (China), DE (Germany), EU (European Union), FR (France), IT (Italy), JP (Japan), MX (Mexico), US (USA), UK (United Kingdom) Code so that your app gracefully handles any future changes to this list.
	Categoryid string `json:"categoryId,omitempty"` // The ID of the leaf category for this item. A leaf category is the lowest level in that category and has no children.
	Image Image `json:"image,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
	Itemaffiliateweburl string `json:"itemAffiliateWebUrl,omitempty"` // The URL of the View Item page of the item, which includes the affiliate tracking ID. This field is only returned if the eBay partner enables affiliate tracking for the item by including the X-EBAY-C-ENDUSERCTX request header in the method. Note: eBay Partner Network, in order to be commissioned for your sales, you must use this URL to forward your buyer to the ebay.com site.
	Legacyitemid string `json:"legacyItemId,omitempty"` // The unique identifier of the eBay listing that contains the item. This is the traditional/legacy ID that is often seen in the URL of the listing View Item page.
	Primaryproductreviewrating ReviewRating `json:"primaryProductReviewRating,omitempty"` // The type that defines the fields for the rating of a product review.
	Selleritemrevision string `json:"sellerItemRevision,omitempty"` // An identifier generated/incremented when a seller revises the item. There are two types of item revisions: Seller changes, such as changing the title eBay system changes, such as changing the quantity when an item is purchased This ID is changed only when the seller makes a change to the item. This means you cannot use this value to determine if the quantity has changed.
	Conditiondescription string `json:"conditionDescription,omitempty"` // A full text description for the condition of the item. This field elaborates on the value specified in the condition field and provides full details for the condition of the item. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Lotsize int `json:"lotSize,omitempty"` // The number of items in a lot. In other words, a lot size is the number of items that are being sold together. A lot is a set of two or more items included in a single listing that must be purchased together in a single order line item. All the items in the lot are the same but there can be multiple items in a single lot, such as the package of batteries shown in the example below. Item Lot Definition Lot Size A package of 24 AA batteries A box of 10 packages 10 A P235/75-15 Goodyear tire 4 tires 4 Fashion Jewelry Rings Package of 100 assorted rings 100 Note: Lots are not supported in all categories.
	Qualifiedprograms []string `json:"qualifiedPrograms,omitempty"` // An array of the qualified programs available for the item, such as EBAY_PLUS. eBay Plus is a premium account option for buyers, which provides benefits such as fast free domestic shipping and free returns on selected items. Top-Rated eBay sellers must opt in to eBay Plus to be able to offer the program on qualifying listings. Sellers must commit to next-day delivery of those items. Note: eBay Plus is available only to buyers in Germany, Austria, and Australia marketplaces.
	Pricedisplaycondition string `json:"priceDisplayCondition,omitempty"` // Indicates when in the buying flow the item's price can appear for minimum advertised price (MAP) items, which is the lowest price a retailer can advertise/show for this item. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceDisplayConditionEnum'>eBay API documentation</a>
	Uniquebiddercount int `json:"uniqueBidderCount,omitempty"` // This integer value indicates the number of different eBay users who have placed one or more bids on an auction item. This field is only applicable to auction items.
	Gtin string `json:"gtin,omitempty"` // The unique Global Trade Item number of the item as defined by https://www.gtin.info. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value.
	Topratedbuyingexperience bool `json:"topRatedBuyingExperience,omitempty"` // This indicates if the item a top-rated plus item. There are three benefits of a top-rated plus item; a minimum 30-day money-back return policy, shipping the items in 1 business day with tracking provided, and the added comfort of knowing this item is from experienced sellers with the highest buyer ratings. See the Top Rated Plus Items and Becoming a Top Rated Seller and qualifying for Top Rated Plus help topics for more information.
	Enabledforguestcheckout bool `json:"enabledForGuestCheckout,omitempty"` // This indicates if the item can be purchased using Guest Checkout in the Order API. You can use this flag to exclude items from your inventory that are not eligible for Guest Checkout, such as gift cards.
	Productficheweburl string `json:"productFicheWebUrl,omitempty"` // The URL of a page containing the manufacturer's specification of this item, which helps buyers make a purchasing decision. This information is available only for items that include the European energy efficiency rating (EEK) but is not available for all items with an EEK rating and is returned only if this information is available. The EEK rating of the item is returned in the energyEfficiencyClass field.
	Additionalimages []Image `json:"additionalImages,omitempty"` // An array of containers with the URLs for the images that are in addition to the primary image. The primary image is returned in the image.imageUrl field.
	Mpn string `json:"mpn,omitempty"` // The manufacturer's part number, which is a unique number that identifies a specific product. To identify the product, this is always used along with brand.
	Material string `json:"material,omitempty"` // (Primary Item Aspect) Text describing what the item is made of. For example, silk. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Inferredepid string `json:"inferredEpid,omitempty"` // The ePID (eBay Product ID of a product from the eBay product catalog) for the item, which has been programmatically determined by eBay using the item's title, aspects, and other data. If the seller provided an ePID for the item, the seller's value is returned in the epid field. Note: This field is returned only for authorized Partners.
	Currentbidprice ConvertedAmount `json:"currentBidPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Subtitle string `json:"subtitle,omitempty"` // A subtitle is optional and allows the seller to provide more information about the product, possibly including keywords that may assist with search results.
	Taxes []Taxes `json:"taxes,omitempty"` // The container for the tax information for the item.
	Title string `json:"title,omitempty"` // The seller-created title of the item. Maximum Length: 80 characters
	Price ConvertedAmount `json:"price,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Unitprice ConvertedAmount `json:"unitPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Condition string `json:"condition,omitempty"` // A short text description for the condition of the item, such as New or Used. For a list of condition names, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
}

// CoreItem represents the CoreItem schema from the OpenAPI specification
type CoreItem struct {
	Primaryitemgroup ItemGroupSummary `json:"primaryItemGroup,omitempty"` // The type that defines the fields for the details of each item in an item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. When an item group is created, one of the item variations, such as the red shirt size L, is chosen as the "parent". All the other items in the group are the children, such as the blue shirt size L, red shirt size M, etc. <br /><br /><span class="tablenote"><b> Note: </b> This container is returned only if the <b> item_id</b> in the request is an item group (parent ID of an item with variations).</span>
	Gender string `json:"gender,omitempty"` // (Primary Item Aspect) The gender for the item. This is used for items that could vary by gender, such as clothing. For example: male, female, or unisex. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Pricedisplaycondition string `json:"priceDisplayCondition,omitempty"` // Indicates when in the buying flow the item's price can appear for minimum advertised price (MAP) items, which is the lowest price a retailer can advertise/show for this item. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceDisplayConditionEnum'>eBay API documentation</a>
	Agegroup string `json:"ageGroup,omitempty"` // (Primary Item Aspect) The age group for which the product is recommended. For example, newborn, infant, toddler, kids, adult, etc. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Enabledforguestcheckout bool `json:"enabledForGuestCheckout,omitempty"` // This indicates if the item can be purchased using Guest Checkout in the Order API. You can use this flag to exclude items from your inventory that are not eligible for Guest Checkout, such as gift cards.
	Sizesystem string `json:"sizeSystem,omitempty"` // (Primary Item Aspect) The sizing system of the country. All the item aspects, including this aspect, are returned in the localizedAspects container. Valid Values: AU (Australia), BR (Brazil), CN (China), DE (Germany), EU (European Union), FR (France), IT (Italy), JP (Japan), MX (Mexico), US (USA), UK (United Kingdom) Code so that your app gracefully handles any future changes to this list.
	Qualifiedprograms []string `json:"qualifiedPrograms,omitempty"` // An array of the qualified programs available for the item, such as EBAY_PLUS. eBay Plus is a premium account option for buyers, which provides benefits such as fast free domestic shipping and free returns on selected items. Top-Rated eBay sellers must opt in to eBay Plus to be able to offer the program on qualifying listings. Sellers must commit to next-day delivery of those items. Note: eBay Plus is available only to buyers in Germany, Austria, and Australia marketplaces.
	Topratedbuyingexperience bool `json:"topRatedBuyingExperience,omitempty"` // This indicates if the item a top-rated plus item. There are three benefits of a top-rated plus item; a minimum 30-day money-back return policy, shipping the items in 1 business day with tracking provided, and the added comfort of knowing this item is from experienced sellers with the highest buyer ratings. See the Top Rated Plus Items and Becoming a Top Rated Seller and qualifying for Top Rated Plus help topics for more information.
	Energyefficiencyclass string `json:"energyEfficiencyClass,omitempty"` // This indicates the European energy efficiency rating (EEK) of the item. This field is returned only if the seller specified the energy efficiency rating. The rating is a set of energy efficiency classes from A to G, where 'A' is the most energy efficient and 'G' is the least efficient. This rating helps buyers choose between various models. When the manufacturer's specifications for this item are available, the link to this information is returned in the productFicheWebUrl field.
	Conditionid string `json:"conditionId,omitempty"` // The identifier of the condition of the item. For example, 1000 is the identifier for NEW. For a list of condition names and IDs, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Size string `json:"size,omitempty"` // (Primary Item Aspect) The size of the item. For example, '7' for a size 7 shoe. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Seller SellerDetail `json:"seller,omitempty"` // The type that defines the fields for basic and detailed information about the seller of the item returned by the <b> item</b> resource.
	Unitpricingmeasure string `json:"unitPricingMeasure,omitempty"` // The designation, such as size, weight, volume, count, etc., that was used to specify the quantity of the item. This helps buyers compare prices. For example, the following tells the buyer that the item is 7.99 per 100 grams. &quot;unitPricingMeasure&quot;: &quot;100g&quot;, &quot;unitPrice&quot;: { &nbsp;&nbsp;&quot;value&quot;: &quot;7.99&quot;, &nbsp;&nbsp;&quot;currency&quot;: &quot;GBP&quot;
	Subtitle string `json:"subtitle,omitempty"` // A subtitle is optional and allows the seller to provide more information about the product, possibly including keywords that may assist with search results.
	Estimatedavailabilities []EstimatedAvailability `json:"estimatedAvailabilities,omitempty"` // The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned.
	Productficheweburl string `json:"productFicheWebUrl,omitempty"` // The URL of a page containing the manufacturer's specification of this item, which helps buyers make a purchasing decision. This information is available only for items that include the European energy efficiency rating (EEK) but is not available for all items with an EEK rating and is returned only if this information is available. The EEK rating of the item is returned in the energyEfficiencyClass field.
	Color string `json:"color,omitempty"` // (Primary Item Aspect) Text describing the color of the item. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Itemenddate string `json:"itemEndDate,omitempty"` // The date and time up to which the items can be purchased. This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: This field is not returned for Good 'Til Cancelled (GTC) listings.
	Itemaffiliateweburl string `json:"itemAffiliateWebUrl,omitempty"` // The URL of the View Item page of the item, which includes the affiliate tracking ID. This field is only returned if the eBay partner enables affiliate tracking for the item by including the X-EBAY-C-ENDUSERCTX request header in the method. Note: eBay Partner Network, in order to be commissioned for your sales, you must use this URL to forward your buyer to the ebay.com site.
	Itemlocation Address `json:"itemLocation,omitempty"` // The type that defines the fields for an address.
	Material string `json:"material,omitempty"` // (Primary Item Aspect) Text describing what the item is made of. For example, silk. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Localizedaspects []TypedNameValue `json:"localizedAspects,omitempty"` // An array of containers that show the complete list of the aspect name/value pairs that describe the variation of the item.
	Eligibleforinlinecheckout bool `json:"eligibleForInlineCheckout,omitempty"` // This field indicates if the item can be purchased using the Buy Order API. If the value of this field is true, this indicates that the item can be purchased using the Order API. If the value of this field is false, this indicates that the item cannot be purchased using the Order API and must be purchased on the eBay site.
	Itemweburl string `json:"itemWebUrl,omitempty"` // The URL of the View Item page of the item. This enables you to include a &quot;Report Item on eBay&quot; link that takes the buyer to the View Item page on eBay. From there they can report any issues regarding this item to eBay.
	Legacyitemid string `json:"legacyItemId,omitempty"` // The unique identifier of the eBay listing that contains the item. This is the traditional/legacy ID that is often seen in the URL of the listing View Item page.
	Epid string `json:"epid,omitempty"` // An EPID is the eBay product identifier of a product from the eBay product catalog. This indicates the product in which the item belongs.
	Inferredepid string `json:"inferredEpid,omitempty"` // The ePID (eBay Product ID of a product from the eBay product catalog) for the item, which has been programmatically determined by eBay using the item's title, aspects, and other data. If the seller provided an ePID for the item, the seller's value is returned in the epid field. Note: This field is returned only for authorized Partners.
	Shiptolocations ShipToLocations `json:"shipToLocations,omitempty"` // The type that defines the fields that include and exclude geographic regions affecting where the item can be shipped. The seller defines these regions when listing the item.
	Description string `json:"description,omitempty"` // The full description of the item that was created by the seller. This can be plain text or rich content and can be very large.
	Price ConvertedAmount `json:"price,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Bidcount int `json:"bidCount,omitempty"` // This integer value indicates the total number of bids that have been placed against an auction item. This field is returned only for auction items.
	Pattern string `json:"pattern,omitempty"` // (Primary Item Aspect) Text describing the pattern used on the item. For example, paisley. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Uniquebiddercount int `json:"uniqueBidderCount,omitempty"` // This integer value indicates the number of different eBay users who have placed one or more bids on an auction item. This field is only applicable to auction items.
	Selleritemrevision string `json:"sellerItemRevision,omitempty"` // An identifier generated/incremented when a seller revises the item. There are two types of item revisions: Seller changes, such as changing the title eBay system changes, such as changing the quantity when an item is purchased This ID is changed only when the seller makes a change to the item. This means you cannot use this value to determine if the quantity has changed.
	Shortdescription string `json:"shortDescription,omitempty"` // This text string is derived from the item condition and the item aspects (such as size, color, capacity, model, brand, etc.).
	Additionalimages []Image `json:"additionalImages,omitempty"` // An array of containers with the URLs for the images that are in addition to the primary image. The primary image is returned in the image.imageUrl field.
	Conditiondescription string `json:"conditionDescription,omitempty"` // A full text description for the condition of the item. This field elaborates on the value specified in the condition field and provides full details for the condition of the item. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Categoryid string `json:"categoryId,omitempty"` // The ID of the leaf category for this item. A leaf category is the lowest level in that category and has no children.
	Product Product `json:"product,omitempty"` // The type that defines the fields for the product information of the item.
	Condition string `json:"condition,omitempty"` // A short text description for the condition of the item, such as New or Used. For a list of condition names, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Brand string `json:"brand,omitempty"` // (Primary Item Aspect) The name brand of the item, such as Nike, Apple, etc. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Quantitylimitperbuyer int `json:"quantityLimitPerBuyer,omitempty"` // The maximum number for a specific item that one buyer can purchase.
	Title string `json:"title,omitempty"` // The seller-created title of the item. Maximum Length: 80 characters
	Availablecoupons []AvailableCoupon `json:"availableCoupons,omitempty"` // A list of available coupons for the item.
	Shippingoptions []ShippingOption `json:"shippingOptions,omitempty"` // An array of shipping options containers that have the details about cost, carrier, etc. of one shipping option.
	Adultonly bool `json:"adultOnly,omitempty"` // This indicates if the item is for adults only. For more information about adult-only items on eBay, see Adult items policy for sellers and Adult-Only items on eBay for buyers.
	Sizetype string `json:"sizeType,omitempty"` // (Primary Item Aspect) Text describing a size group in which the item would be included, such as regular, petite, plus, big-and-tall or maternity. All the item aspects, including this aspect, are returned in the localizedAspects container.
	Gtin string `json:"gtin,omitempty"` // The unique Global Trade Item number of the item as defined by https://www.gtin.info. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value.
	Currentbidprice ConvertedAmount `json:"currentBidPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Unitprice ConvertedAmount `json:"unitPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Buyingoptions []string `json:"buyingOptions,omitempty"` // A comma separated list of all the purchase options available for the item. The values returned are: FIXED_PRICE - Indicates the buyer can purchase the item for a set price using the Buy It Now button. AUCTION - Indicates the buyer can place a bid for the item. After the first bid is placed, this becomes a live auction item and is the only buying option for this item. BEST_OFFER - Indicates the buyer can send the seller a price they're willing to pay for the item. The seller can accept, reject, or send a counter offer. For more information on how this works, see Making a Best Offer. Code so that your app gracefully handles any future changes to this list.
	Itemid string `json:"itemId,omitempty"` // The unique RESTful identifier of the item.
	Returnterms ItemReturnTerms `json:"returnTerms,omitempty"` // The type that defines the fields for the seller's return policy.
	Lotsize int `json:"lotSize,omitempty"` // The number of items in a lot. In other words, a lot size is the number of items that are being sold together. A lot is a set of two or more items included in a single listing that must be purchased together in a single order line item. All the items in the lot are the same but there can be multiple items in a single lot, such as the package of batteries shown in the example below. Item Lot Definition Lot Size A package of 24 AA batteries A box of 10 packages 10 A P235/75-15 Goodyear tire 4 tires 4 Fashion Jewelry Rings Package of 100 assorted rings 100 Note: Lots are not supported in all categories.
	Categorypath string `json:"categoryPath,omitempty"` // Text that shows the category hierarchy of the item. For example: Computers/Tablets &amp; Networking, Laptops &amp; Netbooks, PC Laptops &amp; Netbooks
	Paymentmethods []PaymentMethod `json:"paymentMethods,omitempty"` // The payment methods for the item, including the payment method types, brands, and instructions for the buyer.
	Mpn string `json:"mpn,omitempty"` // The manufacturer's part number, which is a unique number that identifies a specific product. To identify the product, this is always used along with brand.
	Reservepricemet bool `json:"reservePriceMet,omitempty"` // This indicates if the reserve price of the item has been met. A reserve price is set by the seller and is the minimum amount the seller is willing to sell the item for. If the highest bid is not equal to or higher than the reserve price when the auction ends, the listing ends and the item is not sold. Note: This is returned only for auctions that have a reserve price.
	Marketingprice MarketingPrice `json:"marketingPrice,omitempty"` // The type that defines the fields that describe a seller discount.
	Minimumpricetobid ConvertedAmount `json:"minimumPriceToBid,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Taxes []Taxes `json:"taxes,omitempty"` // The container for the tax information for the item.
	Primaryproductreviewrating ReviewRating `json:"primaryProductReviewRating,omitempty"` // The type that defines the fields for the rating of a product review.
	Image Image `json:"image,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
}

// SellerDetail represents the SellerDetail schema from the OpenAPI specification
type SellerDetail struct {
	Sellerlegalinfo SellerLegalInfo `json:"sellerLegalInfo,omitempty"` // The type that defines the fields for the contact information for a seller.
	Username string `json:"username,omitempty"` // The user name created by the seller for use on eBay.
	Feedbackpercentage string `json:"feedbackPercentage,omitempty"` // The percentage of the total positive feedback.
	Feedbackscore int `json:"feedbackScore,omitempty"` // The feedback score of the seller. This value is based on the ratings from eBay members that bought items from this seller.
	Selleraccounttype string `json:"sellerAccountType,omitempty"` // This indicates if the seller is a business or an individual. This is determined when the seller registers with eBay. If they register for a business account, this value will be BUSINESS. If they register for a private account, this value will be INDIVIDUAL. This designation is required by the tax laws in the following countries: This field is returned only on the following sites. EBAY_AT, EBAY_BE, EBAY_CH, EBAY_DE, EBAY_ES, EBAY_FR, EBAY_GB, EBAY_IE, EBAY_IT, EBAY_PL Valid Values: BUSINESS or INDIVIDUAL Code so that your app gracefully handles any future changes to this list.
}

// ShipToRegion represents the ShipToRegion schema from the OpenAPI specification
type ShipToRegion struct {
	Regionid string `json:"regionId,omitempty"` // The unique identifier of the shipping region. The value returned here is dependent on the corresponding regionType value. The regionId value for a region does not vary based on the eBay marketplace. However, the corresponding regionName value for a region is a localized, text-based description of the shipping region. If the regionType value is WORLDWIDE, the regionId value will also be WORLDWIDE. If the regionType value is WORLD_REGION, the regionId value will be one of the following: AFRICA, AMERICAS, ASIA, AUSTRALIA, CENTRAL_AMERICA_AND_CARIBBEAN, EUROPE, EUROPEAN_UNION, GREATER_CHINA, MIDDLE_EAST, NORTH_AMERICA, OCEANIA, SOUTH_AMERICA, SOUTHEAST_ASIA or CHANNEL_ISLANDS. If the regionType value is COUNTRY, the regionId value will be the two-letter code for the country, as defined in the ISO 3166 standard. If the regionType value is STATE_OR_PROVINCE, the regionId value will either be the two-letter code for US states and DC (as defined on this Social Security Administration page), or the two-letter code for Canadian provinces (as defined by this Canada Post page). If the regionType value is COUNTRY_REGION, the regionId value may be one of following: _AH (if a seller is not willing to ship to Alaska/Hawaii), _PR (if the seller is not willing to ship to US Protectorates), _AP (if seller is not willing to ship to a US Army or Fleet Post Office), and PO_BOX (if the seller is not willing to ship to a Post Office Box).
	Regionname string `json:"regionName,omitempty"` // A localized text string that indicates the name of the shipping region. The value returned here is dependent on the corresponding regionType value. If the regionType value is WORLDWIDE, the regionName value will show Worldwide. If the regionType value is WORLD_REGION, the regionName value will be a localized text string for one of the following large geographical regions: Africa, Americas, Asia, Australia, Central America and Caribbean, Europe, European Union, Greater China, Middle East, North America, Oceania, South America, Southeast Asia, or Channel Islands. If the regionType value is COUNTRY, the regionName value will be a localized text string for any country in the world. If the regionType value is STATE_OR_PROVINCE, the regionName value will be a localized text string for any US state or Canadian province. If the regionType value is COUNTRY_REGION, the regionName value may be a localized version of one of the following: Alaska/Hawaii, US Protectorates, APO/FPO (Army or Fleet Post Office), or PO BOX.
	Regiontype string `json:"regionType,omitempty"` // An enumeration value that indicates the level or type of shipping region. Valid Values: COUNTRY_REGION - Indicates the region is a domestic region or special location within a country. STATE_OR_PROVINCE - Indicates the region is a state or province within a country, such as California or New York in the US, or Ontario or Alberta in Canada. COUNTRY - Indicates the region is a single country. WORLD_REGION - Indicates the region is a world region, such as Africa, the Middle East, or Southeast Asia. WORLDWIDE - Indicates the region is the entire world. This value is only applicable for included shiping regions, and not excluded shipping regions. For more detail on the actual regionName/regionId values that will be returned based on the regionType value, see the regionId and/or regionName field descriptions. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:RegionTypeEnum'>eBay API documentation</a>
}

// ConditionDistribution represents the ConditionDistribution schema from the OpenAPI specification
type ConditionDistribution struct {
	Refinementhref string `json:"refinementHref,omitempty"` // The HATEOAS reference of this condition.
	Condition string `json:"condition,omitempty"` // The text describing the condition of the item, such as New or Used. For a list of condition names, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Conditionid string `json:"conditionId,omitempty"` // The identifier of the condition. For example, 1000 is the identifier for NEW. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Matchcount int `json:"matchCount,omitempty"` // The number of items having the condition.
}

// Amount represents the Amount schema from the OpenAPI specification
type Amount struct {
	Currency string `json:"currency,omitempty"` // The list of valid currencies. Each ISO 4217 currency code includes the currency name followed by the numeric value. For example, the Canadian Dollar code (CAD) would take the following form: Canadian Dollar, 124. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
	Value string `json:"value,omitempty"` // The value of the discounted amount.
}

// LegalAddress represents the LegalAddress schema from the OpenAPI specification
type LegalAddress struct {
	Addressline2 string `json:"addressLine2,omitempty"` // The second line of the street address. This field is not always used, but can be used for 'Suite Number' or 'Apt Number'.
	City string `json:"city,omitempty"` // The city of the address.
	Country string `json:"country,omitempty"` // The two-letter ISO 3166 standard of the country of the address. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
	Countryname string `json:"countryName,omitempty"` // The name of the country of the address.
	County string `json:"county,omitempty"` // The name of the county of the address.
	Postalcode string `json:"postalCode,omitempty"` // The postal code of the address.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The state or province of the address.
	Addressline1 string `json:"addressLine1,omitempty"` // The first line of the street address.
}

// ShippingOption represents the ShippingOption schema from the OpenAPI specification
type ShippingOption struct {
	Fulfilledthrough string `json:"fulfilledThrough,omitempty"` // If the item is being shipped by eBay's Global Shipping Program, this field returns GLOBAL_SHIPPING. Otherwise this field is null. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:FulfilledThroughEnum'>eBay API documentation</a>
	Trademarksymbol string `json:"trademarkSymbol,omitempty"` // Any trademark symbol, such as &trade; or &reg;, that needs to be shown in superscript next to the shipping service name.
	Guaranteeddelivery bool `json:"guaranteedDelivery,omitempty"` // Indicates if the seller has committed to shipping the item with eBay Guaranteed Delivery. With eBay Guaranteed Delivery, the seller is committed to getting the line item to the buyer within 4 business days or less. See the Buying items with eBay Guaranteed Delivery help topic for more details about eBay Guaranteed Delivery.
	Importcharges ConvertedAmount `json:"importCharges,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Maxestimateddeliverydate string `json:"maxEstimatedDeliveryDate,omitempty"` // The end date of the delivery window (latest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: For the best accuracy, always include the location of where the item is be shipped in the contextualLocation values of the X-EBAY-C-ENDUSERCTX request header.
	Cutoffdateusedforestimate string `json:"cutOffDateUsedForEstimate,omitempty"` // The deadline date that the item must be purchased by in order to be received by the buyer within the delivery window ( maxEstimatedDeliveryDate and minEstimatedDeliveryDate fields). This field is returned only for items that are eligible for 'Same Day Handling'. For these items, the value of this field is what is displayed in the Delivery line on the View Item page. This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.
	TypeField string `json:"type,omitempty"` // The type of a shipping option, such as EXPEDITED, ONE_DAY, STANDARD, ECONOMY, PICKUP, etc.
	Additionalshippingcostperunit ConvertedAmount `json:"additionalShippingCostPerUnit,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Shippingservicecode string `json:"shippingServiceCode,omitempty"` // The type of shipping service. For example, USPS First Class.
	Shippingcost ConvertedAmount `json:"shippingCost,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Minestimateddeliverydate string `json:"minEstimatedDeliveryDate,omitempty"` // The start date of the delivery window (earliest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. Note: For the best accuracy, always include the location of where the item is be shipped in the contextualLocation values of the X-EBAY-C-ENDUSERCTX request header.
	Shiptolocationusedforestimate ShipToLocation `json:"shipToLocationUsedForEstimate,omitempty"` // The type that defines the fields for the country and postal code of where an item is to be shipped.
	Shippingcosttype string `json:"shippingCostType,omitempty"` // Indicates the class of the shipping cost. Valid Values: FIXED or CALCULATED Code so that your app gracefully handles any future changes to this list.
	Quantityusedforestimate int `json:"quantityUsedForEstimate,omitempty"` // The number of items used when calculating the estimation information.
	Shippingcarriercode string `json:"shippingCarrierCode,omitempty"` // The name of the shipping provider, such as FedEx, or USPS.
}

// ItemSummary represents the ItemSummary schema from the OpenAPI specification
type ItemSummary struct {
	Price ConvertedAmount `json:"price,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Title string `json:"title,omitempty"` // The seller-created title of the item. Maximum Length: 80 characters
	Itemid string `json:"itemId,omitempty"` // The unique RESTful identifier of the item.
	Additionalimages []Image `json:"additionalImages,omitempty"` // An array of containers with the URLs for the images that are in addition to the primary image. The primary image is returned in the image.imageUrl field.
	Pricedisplaycondition string `json:"priceDisplayCondition,omitempty"` // Indicates when in the buying flow the item's price can appear for minimum advertised price (MAP) items, which is the lowest price a retailer can advertise/show for this item. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceDisplayConditionEnum'>eBay API documentation</a>
	Shippingoptions []ShippingOptionSummary `json:"shippingOptions,omitempty"` // This container returns the shipping options available to ship the item.
	Shortdescription string `json:"shortDescription,omitempty"` // This text string is derived from the item condition and the item aspects (such as size, color, capacity, model, brand, etc.). Sometimes the title doesn't give enough information but the description is too big. Surfacing the shortDescription can often provide buyers with the additional information that could help them make a buying decision. For example: &quot; title&quot;: &quot;Petrel U42W FPV Drone RC Quadcopter w/HD Camera Live Video One Key Off / Landing&quot;, &quot;shortDescription&quot;: &quot;1 U42W Quadcopter. Syma X5SW-V3 Wifi FPV RC Drone Quadcopter 2.4Ghz 6-Axis Gyro with Headless Mode. Syma X20 Pocket Drone 2.4Ghz Mini RC Quadcopter Headless Mode Altitude Hold. One Key Take Off / Landing function: allow beginner to easy to fly the drone without any skill.&quot;, Restriction: This field is returned by the search method only when fieldgroups = EXTENDED.
	Categories []Category `json:"categories,omitempty"` // This container returns the primary category ID of the item (as well as the secondary category if the item was listed in two categories).
	Energyefficiencyclass string `json:"energyEfficiencyClass,omitempty"` // This indicates the European energy efficiency rating (EEK) of the item. Energy efficiency ratings apply to products listed by commercial vendors in electronics categories only. Currently, this field is only applicable for the Germany site, and is only returned if the seller specified the energy efficiency rating through item specifics at listing time. Rating values include A+++, A++, A+, A, B, C, D, E, F, and G.
	Availablecoupons bool `json:"availableCoupons,omitempty"` // This boolean attribute indicates if coupons are available for the item.
	Itemgrouptype string `json:"itemGroupType,omitempty"` // The indicates the item group type. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. Currently only the SELLER_DEFINED_VARIATIONS is supported and indicates this is an item group created by the seller. Note: This field is returned only for item groups. Code so that your app gracefully handles any future changes to this list.
	Itemhref string `json:"itemHref,omitempty"` // The URI for the Browse API getItem method, which can be used to retrieve more details about items in the search results.
	Marketingprice MarketingPrice `json:"marketingPrice,omitempty"` // The type that defines the fields that describe a seller discount.
	Buyingoptions []string `json:"buyingOptions,omitempty"` // A comma separated list of all the purchase options available for the item. Values Returned: FIXED_PRICE - Indicates the buyer can purchase the item for a set price using the Buy It Now button. AUCTION - Indicates the buyer can place a bid for the item. After the first bid is placed, this becomes a live auction item and is the only buying option for this item. BEST_OFFER - Items where the buyer can send the seller a price they're willing to pay for the item. The seller can accept, reject, or send a counter offer. For details about Best Offer, see Best Offer. Code so that your app gracefully handles any future changes to this list.
	Pickupoptions []PickupOptionSummary `json:"pickupOptions,omitempty"` // This container returns the local pickup options available to the buyer. This container is only returned if the user is searching for local pickup items and set the local pickup filters in the method request.
	Legacyitemid string `json:"legacyItemId,omitempty"` // The unique identifier of the eBay listing that contains the item. This is the traditional/legacy ID that is often seen in the URL of the listing View Item page.
	Unitpricingmeasure string `json:"unitPricingMeasure,omitempty"` // The designation, such as size, weight, volume, count, etc., that was used to specify the quantity of the item. This helps buyers compare prices. For example, the following tells the buyer that the item is 7.99 per 100 grams. &quot;unitPricingMeasure&quot;: &quot;100g&quot;, &quot;unitPrice&quot;: { &nbsp;&nbsp;&quot;value&quot;: &quot;7.99&quot;, &nbsp;&nbsp;&quot;currency&quot;: &quot;GBP&quot;
	Image Image `json:"image,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
	Compatibilitymatch string `json:"compatibilityMatch,omitempty"` // This indicates how well the item matches the compatibility_filter product attributes. Valid Values: EXACT or POSSIBLE Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityMatchEnum'>eBay API documentation</a>
	Unitprice ConvertedAmount `json:"unitPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Itemweburl string `json:"itemWebUrl,omitempty"` // The URL to the View Item page of the item. This enables you to include a &quot;Report Item on eBay&quot; hyperlink that takes the buyer to the View Item page on eBay. From there they can report any issues regarding this item to eBay.
	Condition string `json:"condition,omitempty"` // The text describing the condition of the item, such as New or Used. For a list of condition names, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Itemgrouphref string `json:"itemGroupHref,omitempty"` // The HATEOAS reference of the parent page of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. Note: This field is returned only for item groups.
	Qualifiedprograms []string `json:"qualifiedPrograms,omitempty"` // An array of the qualified programs available for the item, such as EBAY_PLUS. eBay Plus is a premium account option for buyers, which provides benefits such as fast free domestic shipping and free returns on selected items. Top-Rated eBay sellers must opt in to eBay Plus to be able to offer the program on qualifying listings. Sellers must commit to next-day delivery of those items. Note: eBay Plus is available only to buyers in Germany, Austria, and Australia marketplaces.
	Thumbnailimages []Image `json:"thumbnailImages,omitempty"` // An array of thumbnail images for the item.
	Currentbidprice ConvertedAmount `json:"currentBidPrice,omitempty"` // This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
	Distancefrompickuplocation TargetLocation `json:"distanceFromPickupLocation,omitempty"` // The type that defines the fields for the distance between the item location and the buyer's location.
	Bidcount int `json:"bidCount,omitempty"` // This integer value indicates the total number of bids that have been placed for an auction item. This field is only returned for auction items.
	Seller Seller `json:"seller,omitempty"` // The type that defines the fields for basic information about the seller of the item returned by the <b> item_summary</b> resource.
	Itemlocation ItemLocationImpl `json:"itemLocation,omitempty"` // The type that defines the fields for the location of an item, such as information typically used for an address, including postal code, county, state/province, street address, city, and country (2-digit ISO code).
	Conditionid string `json:"conditionId,omitempty"` // The identifier of the condition of the item. For example, 1000 is the identifier for NEW. For a list of condition names and IDs, see Item Condition IDs and Names. Code so that your app gracefully handles any future changes to this list. Note: In the US and Australian marketplaces, Condition ID 2000 now maps to an item condition of 'Certified Refurbished', but this item condition is only available for use for a select number of US and Australian sellers. For all other marketplaces besides the US and Australia, Condition ID 2000 still maps to 'Manufacturer Refurbished'.
	Compatibilityproperties []CompatibilityProperty `json:"compatibilityProperties,omitempty"` // This container returns only the product attributes that are compatible with the item. These attributes were specified in the compatibility_filter in the request. This means that if you passed in 5 attributes and only 4 are compatible, only those 4 are returned. If none of the attributes are compatible, this container is not returned.
	Epid string `json:"epid,omitempty"` // An ePID is the eBay product identifier of a product from the eBay product catalog. This indicates the product in which the item belongs.
	Adultonly bool `json:"adultOnly,omitempty"` // This indicates if the item is for adults only. For more information about adult-only items on eBay, see Adult items policy for sellers and Adult-Only items on eBay for buyers.
	Itemaffiliateweburl string `json:"itemAffiliateWebUrl,omitempty"` // The URL to the View Item page of the item, which includes the affiliate tracking ID. This field is only returned if the seller enables affiliate tracking for the item by including the X-EBAY-C-ENDUSERCTX request header in the method. Note: eBay Partner Network, in order to receive a commission for your sales, you must use this URL to forward your buyer to the ebay.com site.
}

// ItemLocationImpl represents the ItemLocationImpl schema from the OpenAPI specification
type ItemLocationImpl struct {
	Addressline1 string `json:"addressLine1,omitempty"` // The first line of the street address.
	Addressline2 string `json:"addressLine2,omitempty"` // The second line of the street address. This field may contain such values as an apartment or suite number.
	City string `json:"city,omitempty"` // The city in which the item is located. Restriction: This field is populated in the search method response only when fieldgroups = EXTENDED.
	Country string `json:"country,omitempty"` // The two-letter ISO 3166 standard code that indicates the country in which the item is located. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
	County string `json:"county,omitempty"` // The county in which the item is located.
	Postalcode string `json:"postalCode,omitempty"` // The postal code (or zip code in US) where the item is located. Sellers set a postal code (or zip code in US) for items when they are listed. The postal code is used for calculating proximity searches. It is anonymized when returned in itemLocation.postalCode via the API.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The state or province in which the item is located.
}

// TargetLocation represents the TargetLocation schema from the OpenAPI specification
type TargetLocation struct {
	Value string `json:"value,omitempty"` // This value indicates the distance (measured in the measurement unit in the unitOfMeasure field) between the item location and the buyer's location.
	Unitofmeasure string `json:"unitOfMeasure,omitempty"` // This value shows the unit of measurement used to measure the distance between the location of the item and the buyer's location. This value is typically mi or km.
}

// PickupOptionSummary represents the PickupOptionSummary schema from the OpenAPI specification
type PickupOptionSummary struct {
	Pickuplocationtype string `json:"pickupLocationType,omitempty"` // This container returns the local pickup options available to the buyer. Possible values are ARRANGED_LOCATION and STORE.
}

// AvailableCoupon represents the AvailableCoupon schema from the OpenAPI specification
type AvailableCoupon struct {
	Message string `json:"message,omitempty"` // A description of the coupon. Note: The value returned in the termsWebUrl field should appear for all experiences when displaying coupons. The value in the availableCoupons.message field must also be included, if returned in the API response.
	Redemptioncode string `json:"redemptionCode,omitempty"` // The coupon code.
	Termsweburl string `json:"termsWebUrl,omitempty"` // The URL to the coupon terms of use. Note: The value returned in the termsWebUrl field should appear for all experiences when displaying coupons. The value in the availableCoupons.message field must also be included, if returned in the API response.
	Constraint string `json:"constraint,omitempty"` // The limitations or restrictions of the coupon. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:CouponConstraint'>eBay API documentation</a>
	Discountamount Amount `json:"discountAmount,omitempty"`
	Discounttype string `json:"discountType,omitempty"` // The type of discount that the coupon applies. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:CouponDiscountType'>eBay API documentation</a>
}

// Aspect represents the Aspect schema from the OpenAPI specification
type Aspect struct {
	Localizedname string `json:"localizedName,omitempty"` // The text representing the name of the aspect for the name/value pair, such as Brand.
	Localizedvalues []string `json:"localizedValues,omitempty"` // The text representing the value of the aspect for the name/value pair, such as Apple.
}

// PaymentMethodBrand represents the PaymentMethodBrand schema from the OpenAPI specification
type PaymentMethodBrand struct {
	Logoimage Image `json:"logoImage,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
	Paymentmethodbrandtype string `json:"paymentMethodBrandType,omitempty"` // The payment method brand, such as Visa or PayPal. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PaymentMethodBrandEnum'>eBay API documentation</a>
}

// ItemReturnTerms represents the ItemReturnTerms schema from the OpenAPI specification
type ItemReturnTerms struct {
	Returnsaccepted bool `json:"returnsAccepted,omitempty"` // Indicates whether the seller accepts returns for the item.
	Extendedholidayreturnsoffered bool `json:"extendedHolidayReturnsOffered,omitempty"` // This indicates if the seller has enabled the Extended Holiday Returns feature on the item. Extended Holiday Returns are only applicable during the US holiday season, and gives buyers extra time to return an item. This 'extra time' will typically extend beyond what is set through the returnPeriod value.
	Refundmethod string `json:"refundMethod,omitempty"` // An enumeration value that indicates how a buyer is refunded when an item is returned. Valid Values: MONEY_BACK or MERCHANDISE_CREDIT Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:RefundMethodEnum'>eBay API documentation</a>
	Restockingfeepercentage string `json:"restockingFeePercentage,omitempty"` // This string field indicates the restocking fee percentage that the seller has set on the item. Sellers have the option of setting no restocking fee for an item, or they can set the percentage to 10, 15, or 20 percent. So, if the cost of the item was $100, and the restocking percentage was 20 percent, the buyer would be charged $20 to return that item, so instead of receiving a $100 refund, they would receive $80 due to the restocking fee.
	Returninstructions string `json:"returnInstructions,omitempty"` // Text written by the seller describing what the buyer needs to do in order to return the item.
	Returnmethod string `json:"returnMethod,omitempty"` // An enumeration value that indicates the alternative methods for a full refund when an item is returned. This field is returned if the seller offers the buyer an item replacement or exchange instead of a monetary refund. Valid Values: REPLACEMENT - Indicates that the buyer has the option of receiving money back for the returned item, or they can choose to have the seller replace the item with an identical item. EXCHANGE - Indicates that the buyer has the option of receiving money back for the returned item, or they can exchange the item for another similar item. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:ReturnMethodEnum'>eBay API documentation</a>
	Returnperiod TimeDuration `json:"returnPeriod,omitempty"` // The type that defines the fields for a period of time in the time-measurement units supplied.
	Returnshippingcostpayer string `json:"returnShippingCostPayer,omitempty"` // This enumeration value indicates whether the buyer or seller is responsible for return shipping costs when an item is returned. Valid Values: SELLER - Indicates the seller will pay for the shipping costs to return the item. BUYER - Indicates the buyer will pay for the shipping costs to return the item. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:ReturnShippingCostPayerEnum'>eBay API documentation</a>
}

// CategoryDistribution represents the CategoryDistribution schema from the OpenAPI specification
type CategoryDistribution struct {
	Categoryid string `json:"categoryId,omitempty"` // The identifier of the category.
	Categoryname string `json:"categoryName,omitempty"` // The name of the category, such as Baby &amp; Toddler Clothing.
	Matchcount int `json:"matchCount,omitempty"` // The number of items in this category.
	Refinementhref string `json:"refinementHref,omitempty"` // The HATEOAS reference of this category.
}

// BuyingOptionDistribution represents the BuyingOptionDistribution schema from the OpenAPI specification
type BuyingOptionDistribution struct {
	Refinementhref string `json:"refinementHref,omitempty"` // The HATEOAS reference for this buying option.
	Buyingoption string `json:"buyingOption,omitempty"` // The container that returns the buying option type. This will be AUCTION or FIXED_PRICE or both. For details, see buyingOptions.
	Matchcount int `json:"matchCount,omitempty"` // The number of items having this buying option.
}

// AspectGroup represents the AspectGroup schema from the OpenAPI specification
type AspectGroup struct {
	Aspects []Aspect `json:"aspects,omitempty"` // An array of the name/value pairs for the aspects of the product. For example: BRAND/Apple
	Localizedgroupname string `json:"localizedGroupName,omitempty"` // The name of a group of aspects. In the following example, Product Identifiers and Process are product aspect group names. Under the group name are the product aspect name/value pairs. Product Identifiers &nbsp;&nbsp;&nbsp;Brand/Apple &nbsp;&nbsp;&nbsp;Product Family/iMac Processor &nbsp;&nbsp;&nbsp;Processor Type/Intel &nbsp;&nbsp;&nbsp;Processor Speed/3.10
}

// ErrorParameter represents the ErrorParameter schema from the OpenAPI specification
type ErrorParameter struct {
	Name string `json:"name,omitempty"` // This is the name of input field that caused an issue with the call request.
	Value string `json:"value,omitempty"` // This is the actual value that was passed in for the element specified in the name field.
}

// Items represents the Items schema from the OpenAPI specification
type Items struct {
	Warnings []Error `json:"warnings,omitempty"` // An array of warning messages. These types of errors do not prevent the method from executing but should be checked.
	Items []CoreItem `json:"items,omitempty"` // An arraylist of all the items.
	Total int `json:"total,omitempty"` // The total number of items retrieved.
}

// CompatibilityPayload represents the CompatibilityPayload schema from the OpenAPI specification
type CompatibilityPayload struct {
	Compatibilityproperties []AttributeNameValue `json:"compatibilityProperties,omitempty"` // An array of attribute name/value pairs used to define a specific product. For example: If you wanted to specify a specific car, one of the name/value pairs would be &quot;name&quot; : &quot;Year&quot;, &quot;value&quot; : &quot;2019&quot; For a list of the attributes required for cars and trucks and motorcycles see Check compatibility in the Buy Integration Guide.
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Postalcode string `json:"postalCode,omitempty"` // The postal code (or zip code in US) code of the address. Sellers set a postal code (or zip code in US) for items when they are listed. The postal code is used for calculating proximity searches. It is anonymized when returned in itemLocation.postalCode via the API.
	Stateorprovince string `json:"stateOrProvince,omitempty"` // The state or province of the address. Note: This is conditionally returned in the itemLocation field.
	Addressline1 string `json:"addressLine1,omitempty"` // The first line of the street address. Note: This is conditionally returned in the itemLocation field.
	Addressline2 string `json:"addressLine2,omitempty"` // The second line of the street address. This field is not always used, but can be used for 'Suite Number' or 'Apt Number'.
	City string `json:"city,omitempty"` // The city of the address.
	Country string `json:"country,omitempty"` // The two-letter ISO 3166 standard of the country of the address. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
	County string `json:"county,omitempty"` // The county of the address.
}

// ItemGroupSummary represents the ItemGroupSummary schema from the OpenAPI specification
type ItemGroupSummary struct {
	Itemgrouptitle string `json:"itemGroupTitle,omitempty"` // The title of the item that appears on the item group page. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.
	Itemgrouptype string `json:"itemGroupType,omitempty"` // An enumeration value that indicates the type of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:ItemGroupTypeEnum'>eBay API documentation</a>
	Itemgroupadditionalimages []Image `json:"itemGroupAdditionalImages,omitempty"` // An array of containers with the URLs for images that are in addition to the primary image of the item group. The primary image is returned in the itemGroupImage field.
	Itemgrouphref string `json:"itemGroupHref,omitempty"` // The HATEOAS reference of the parent page of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.
	Itemgroupid string `json:"itemGroupId,omitempty"` // The unique identifier for the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc.
	Itemgroupimage Image `json:"itemGroupImage,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
}

// SearchByImageRequest represents the SearchByImageRequest schema from the OpenAPI specification
type SearchByImageRequest struct {
	Image string `json:"image,omitempty"` // The Base64 string of the image.
}

// ShipToLocation represents the ShipToLocation schema from the OpenAPI specification
type ShipToLocation struct {
	Country string `json:"country,omitempty"` // The two-letter ISO 3166 standard of the country for where the item is to be shipped. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
	Postalcode string `json:"postalCode,omitempty"` // The zip code (postal code) for where the item is to be shipped.
}

// ShipToLocations represents the ShipToLocations schema from the OpenAPI specification
type ShipToLocations struct {
	Regionexcluded []ShipToRegion `json:"regionExcluded,omitempty"` // An array of containers that express the large geographical regions, countries, state/provinces, or special locations within a country where the seller is not willing to ship to.
	Regionincluded []ShipToRegion `json:"regionIncluded,omitempty"` // An array of containers that express the large geographical regions, countries, or state/provinces within a country where the seller is willing to ship to. Prospective buyers must look at the shipping regions under this container, as well as the shipping regions that are under the regionExcluded to see where the seller is willing to ship items. Sellers can specify that they ship 'Worldwide', but then add several large geographical regions (e.g. Asia, Oceania, Middle East) to the exclusion list, or sellers can specify that they ship to Europe and Africa, but then add several individual countries to the exclusion list.
}

// RemoteShopcartResponse represents the RemoteShopcartResponse schema from the OpenAPI specification
type RemoteShopcartResponse struct {
	Unavailablecartitems []CartItem `json:"unavailableCartItems,omitempty"` // An array of items in the cart that are unavailable. This can be for a variety of reasons such as, when the listing has ended or the item is out of stock. Because a cart never expires, these items will remain in the cart until they are removed.
	Warnings []Error `json:"warnings,omitempty"` // An array of warning messages. These type of errors do not prevent the call from executing but should be checked.
	Cartitems []CartItem `json:"cartItems,omitempty"` // An array of the items in the member's eBay cart.
	Cartsubtotal Amount `json:"cartSubtotal,omitempty"`
	Cartweburl string `json:"cartWebUrl,omitempty"` // The URL of the member's eBay cart.
}

// Refinement represents the Refinement schema from the OpenAPI specification
type Refinement struct {
	Buyingoptiondistributions []BuyingOptionDistribution `json:"buyingOptionDistributions,omitempty"` // An array of containers for the all the buying option refinements.
	Categorydistributions []CategoryDistribution `json:"categoryDistributions,omitempty"` // An array of containers for the all the category refinements.
	Conditiondistributions []ConditionDistribution `json:"conditionDistributions,omitempty"` // An array of containers for the all the condition refinements.
	Dominantcategoryid string `json:"dominantCategoryId,omitempty"` // The identifier of the category that most of the items are part of.
	Aspectdistributions []AspectDistribution `json:"aspectDistributions,omitempty"` // An array of containers for the all the aspect refinements.
}

// SellerLegalInfo represents the SellerLegalInfo schema from the OpenAPI specification
type SellerLegalInfo struct {
	Legalcontactlastname string `json:"legalContactLastName,omitempty"` // The seller's last name.
	Termsofservice string `json:"termsOfService,omitempty"` // This is a free-form string created by the seller. This is the seller's terms or condition, which is in addition to the seller's return policies.
	Imprint string `json:"imprint,omitempty"` // This is a free-form string created by the seller. This is information often found on business cards, such as address. This is information used by some countries.
	Legalcontactfirstname string `json:"legalContactFirstName,omitempty"` // The seller's first name.
	Name string `json:"name,omitempty"` // The name of the seller's business.
	Registrationnumber string `json:"registrationNumber,omitempty"` // The seller's registration number. This is information used by some countries.
	Vatdetails []VatDetail `json:"vatDetails,omitempty"` // An array of the seller's VAT (value added tax) IDs and the issuing country. VAT is a tax added by some European countries.
	Phone string `json:"phone,omitempty"` // The seller's business phone number.
	Sellerprovidedlegaladdress LegalAddress `json:"sellerProvidedLegalAddress,omitempty"` // Type that defines the fields for the seller's address.
	Email string `json:"email,omitempty"` // The seller's business email address.
	Fax string `json:"fax,omitempty"` // The seller' business fax number.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Category string `json:"category,omitempty"` // This string value indicates the error category. There are three categories of errors: request errors, application errors, and system errors.
	Domain string `json:"domain,omitempty"` // The name of the primary system where the error occurred. This is relevant for application errors.
	Longmessage string `json:"longMessage,omitempty"` // A detailed description of the condition that caused the error or warning, and information on what to do to correct the problem.
	Message string `json:"message,omitempty"` // A description of the condition that caused the error or warning.
	Outputrefids []string `json:"outputRefIds,omitempty"` // An array of reference IDs that identify the specific response elements most closely associated to the error or warning, if any.
	Parameters []ErrorParameter `json:"parameters,omitempty"` // An array of warning and error messages that return one or more variables contextual information about the error or warning. This is often the field or value that triggered the error or warning.
	Subdomain string `json:"subdomain,omitempty"` // The name of the subdomain in which the error or warning occurred.
	Errorid int `json:"errorId,omitempty"` // A unique code that identifies the particular error or warning that occurred. Your application can use error codes as identifiers in your customized error-handling algorithms.
	Inputrefids []string `json:"inputRefIds,omitempty"` // An array of reference IDs that identify the specific request elements most closely associated to the error or warning, if any.
}

// AspectDistribution represents the AspectDistribution schema from the OpenAPI specification
type AspectDistribution struct {
	Localizedaspectname string `json:"localizedAspectName,omitempty"` // The name of an aspect, such as Brand, Color, etc.
	Aspectvaluedistributions []AspectValueDistribution `json:"aspectValueDistributions,omitempty"` // An array of containers for the various values of the aspect and the match count and a HATEOAS reference ( refinementHref) for this aspect.
}

// CartItem represents the CartItem schema from the OpenAPI specification
type CartItem struct {
	Image Image `json:"image,omitempty"` // Type the defines the details of an image, such as size and image URL. Currently, only <b> imageUrl</b> is populated. The <b> height</b> and <b> width</b> were added for future use.
	Itemid string `json:"itemId,omitempty"` // The RESTful identifier of the item. This identifier is generated when the item was listed. RESTful Item ID Format: v1|#|# For example: v1|272394640372|0 v1|162846450672|461882996982
	Itemweburl string `json:"itemWebUrl,omitempty"` // The URL of the eBay view item page for the item.
	Price Price `json:"price,omitempty"` // The type that defines the fields for the monetary value and currency of the price of the item.
	Quantity int `json:"quantity,omitempty"` // The number of this item the buyer wants to purchase.
	Title string `json:"title,omitempty"` // The title of the item. This can be written by the seller or come from the eBay product catalog.
	Cartitemid string `json:"cartItemId,omitempty"` // The identifier for the item being added to the cart. This is generated when the item is added to the cart.
	Cartitemsubtotal Amount `json:"cartItemSubtotal,omitempty"`
}

// AspectValueDistribution represents the AspectValueDistribution schema from the OpenAPI specification
type AspectValueDistribution struct {
	Localizedaspectvalue string `json:"localizedAspectValue,omitempty"` // The value of an aspect. For example, Red is a value for the aspect Color.
	Matchcount int `json:"matchCount,omitempty"` // The number of items with this aspect.
	Refinementhref string `json:"refinementHref,omitempty"` // A HATEOAS reference for this aspect.
}

// EstimatedAvailability represents the EstimatedAvailability schema from the OpenAPI specification
type EstimatedAvailability struct {
	Availabilitythreshold int `json:"availabilityThreshold,omitempty"` // This field is return only when the seller sets their 'display item quantity' preference to Display &quot;More than 10 available&quot; in your listing (if applicable). The value of this field will be &quot;10&quot;, which is the threshold value. Code so that your app gracefully handles any future changes to this value.
	Availabilitythresholdtype string `json:"availabilityThresholdType,omitempty"` // This field is return only when the seller sets their Display Item Quantity preference to Display &quot;More than 10 available&quot; in your listing (if applicable). The value of this field will be MORE_THAN. This indicates that the seller has more than the 'quantity display preference', which is 10, in stock for this item. The following are the display item quantity preferences the seller can set. Display &quot;More than 10 available&quot; in your listing (if applicable) If the seller enables this preference, this field is returned as long as there are more than 10 of this item in inventory. If the quantity is equal to 10 or drops below 10, this field is not returned and the estimated quantity of the item is returned in the estimatedAvailableQuantity field. Display the exact quantity in your items If the seller enables this preference, the availabilityThresholdType and availabilityThreshold fields are not returned and the estimated quantity of the item is returned in the estimatedAvailableQuantity field. Note: Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. Code so that your app gracefully handles any future changes to these preferences. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:AvailabilityThresholdEnum'>eBay API documentation</a>
	Deliveryoptions []string `json:"deliveryOptions,omitempty"` // An array of available delivery options. Valid Values: SHIP_TO_HOME, SELLER_ARRANGED_LOCAL_PICKUP, IN_STORE_PICKUP, PICKUP_DROP_OFF, or DIGITAL_DELIVERY Code so that your app gracefully handles any future changes to this list.
	Estimatedavailabilitystatus string `json:"estimatedAvailabilityStatus,omitempty"` // An enumeration value representing the inventory status of this item. Note: Be sure to review the itemEndDate field to determine whether the item is available for purchase. Valid Values: IN_STOCK, LIMITED_STOCK, or OUT_OF_STOCK Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:AvailabilityStatusEnum'>eBay API documentation</a>
	Estimatedavailablequantity int `json:"estimatedAvailableQuantity,omitempty"` // The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned.
	Estimatedsoldquantity int `json:"estimatedSoldQuantity,omitempty"` // The estimated number of this item that have been sold.
}

// TypedNameValue represents the TypedNameValue schema from the OpenAPI specification
type TypedNameValue struct {
	Name string `json:"name,omitempty"` // The text representing the name of the aspect for the name/value pair, such as Color.
	TypeField string `json:"type,omitempty"` // This indicates if the value being returned is a string or an array of values. Valid Values: STRING - Indicates the value returned is a string. STRING_ARRAY - Indicates the value returned is an array of strings. Code so that your app gracefully handles any future changes to this list. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:ValueTypeEnum'>eBay API documentation</a>
	Value string `json:"value,omitempty"` // The value of the aspect for the name/value pair, such as Red.
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Imageurl string `json:"imageUrl,omitempty"` // The URL of the image.
	Width int `json:"width,omitempty"` // Reserved for future use.
	Height int `json:"height,omitempty"` // Reserved for future use.
}

// AutoCorrections represents the AutoCorrections schema from the OpenAPI specification
type AutoCorrections struct {
	Q string `json:"q,omitempty"` // The automatically spell-corrected keyword from the request.
}

// CommonDescriptions represents the CommonDescriptions schema from the OpenAPI specification
type CommonDescriptions struct {
	Description string `json:"description,omitempty"` // The item description that is used by more than one of the item variations.
	Itemids []string `json:"itemIds,omitempty"` // A list of item ids that have this description.
}
