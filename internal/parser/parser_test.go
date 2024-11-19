package parser

import (
	"io"
	"slices"
	"strings"
	"testing"
)

func TestFiltersUrls(t *testing.T) {
	URL := "https://test.com/assets/style.css"
	parser := Parser{&ParserConfig{FilteredFileExtensions: []string{".css"}}}
	valid, err := parser.FiltersUrls(&URL)
	if err != nil {
		t.Fatal("error")
	}
	if !valid {
		t.Fatal("URL should be valid because of .css")
	}
}
func TestGetHrefs(t *testing.T) {
	raw_html := `<!doctype html>
<html lang="fr">

<head>
    
        
  <meta charset="utf-8">


  <meta http-equiv="x-ua-compatible" content="ie=edge">



  


      <!-- Google Tag Manager -->
    <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
      new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
              j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
              'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
              })(window,document,'script','dataLayer','GTM-WG4PNKHJ');</script>
    <!-- End Google Tag Manager -->
  
  



  <title>Téléobjectif RedCat 91 WIFD William Optics</title>
  
    
  
  
    
  
  <meta name="description" content="Le Téléobjectif RedCat 91 WIFD William Optics est le dernier né de la fabuleuse série de téléobjectifs RedCat dédiés à l&#039;astrophotographie à grand champ.">
  <meta name="keywords" content="">
    
      <link rel="canonical" href="https://www.astronome.fr/lunettes-astronomiques/3606-teleobjectif-redcat-91-wifd-william-optics.html">
    

  
      

  
    <script type="application/ld+json">
  {
    "@context": "https://schema.org",
    "@id": "#store-organization",
    "@type": "Organization",
    "name" : "L&#039;Astronome",
    "url" : "https://www.astronome.fr/",
  
      "logo": {
        "@type": "ImageObject",
        "url":"https://www.astronome.fr/img/logo-1685139979.jpg"
      }
      }
</script>

<script type="application/ld+json">
  {
    "@context": "https://schema.org",
    "@type": "WebPage",
    "isPartOf": {
      "@type": "WebSite",
      "url":  "https://www.astronome.fr/",
      "name": "L&#039;Astronome"
    },
    "name": "Téléobjectif RedCat 91 WIFD William Optics",
    "url":  "https://www.astronome.fr/lunettes-astronomiques/3606-teleobjectif-redcat-91-wifd-william-optics.html"
  }
</script>


  <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "BreadcrumbList",
      "itemListElement": [
                  {
            "@type": "ListItem",
            "position": 1,
            "name": "Accueil",
            "item": "https://www.astronome.fr/"
          },              {
            "@type": "ListItem",
            "position": 2,
            "name": "Lunettes astronomiques",
            "item": "https://www.astronome.fr/12-lunettes-astronomiques"
          },              {
            "@type": "ListItem",
            "position": 3,
            "name": "Téléobjectif RedCat 91 WIFD William Optics",
            "item": "https://www.astronome.fr/lunettes-astronomiques/3606-teleobjectif-redcat-91-wifd-william-optics.html"
          }          ]
    }
  </script>


  

  
        <script type="application/ld+json">
  {
    "@context": "https://schema.org/",
    "@type": "Product",
    "@id": "#product-snippet-id",
    "name": "Téléobjectif RedCat 91 WIFD William Optics",
    "description": "Le Téléobjectif RedCat 91 WIFD William Optics est le dernier né de la fabuleuse série de téléobjectifs RedCat dédiés à l&#039;astrophotographie à grand champ.",
    "category": "Lunettes astronomiques",
    "image" :"https://www.astronome.fr/11629-home_default/teleobjectif-redcat-91-wifd-william-optics.jpg",    "sku": "T-C-91RDII",
    "mpn": "T-C-91RDII"
        ,
    "brand": {
      "@type": "Brand",
      "name": "WILLIAM OPTICS"
    }
            ,
    "weight": {
        "@context": "https://schema.org",
        "@type": "QuantitativeValue",
        "value": "25.000000",
        "unitCode": "kg"
    }
        ,
    "offers": {
      "@type": "Offer",
      "priceCurrency": "EUR",
      "name": "Téléobjectif RedCat 91 WIFD William Optics",
      "price": "3299",
      "url": "https://www.astronome.fr/lunettes-astronomiques/3606-teleobjectif-redcat-91-wifd-william-optics.html",
      "priceValidUntil": "2024-12-03",
              "image": ["https://www.astronome.fr/11629-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg","https://www.astronome.fr/11630-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg","https://www.astronome.fr/11632-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg","https://www.astronome.fr/11631-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg","https://www.astronome.fr/11633-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg"],
            "sku": "T-C-91RDII",
      "mpn": "T-C-91RDII",
                    "availability": "https://schema.org/InStock",
      "seller": {
        "@type": "Organization",
        "name": "L&#039;Astronome"
      }
    }
      }
</script>


  
    
  



    <meta property="og:type" content="product">
    <meta property="og:url" content="https://www.astronome.fr/lunettes-astronomiques/3606-teleobjectif-redcat-91-wifd-william-optics.html">
    <meta property="og:title" content="Téléobjectif RedCat 91 WIFD William Optics">
    <meta property="og:site_name" content="L&#039;Astronome">
    <meta property="og:description" content="Le Téléobjectif RedCat 91 WIFD William Optics est le dernier né de la fabuleuse série de téléobjectifs RedCat dédiés à l&#039;astrophotographie à grand champ.">
            <meta property="og:image" content="https://www.astronome.fr/11629-thickbox_default/teleobjectif-redcat-91-wifd-william-optics.jpg">
        <meta property="og:image:width" content="1100">
        <meta property="og:image:height" content="1422">
    




      <meta name="viewport" content="width=device-width, initial-scale=1">
  


  <meta name="theme-color" content="#1e90ff">
  <meta name="msapplication-navbutton-color" content="#1e90ff">


  <link rel="icon" type="image/vnd.microsoft.icon" href="https://www.astronome.fr/img/favicon.ico?1685139979">
  <link rel="shortcut icon" type="image/x-icon" href="https://www.astronome.fr/img/favicon.ico?1685139979">
      <link rel="apple-touch-icon" href="/img/cms/logo-1684153297.png">
        <link rel="icon" sizes="192x192" href="/img/cms/logo-1684153297.png">
  




    <link rel="stylesheet" href="https://www.astronome.fr/themes/child_warehouse/assets/cache/theme-f7d352620.css" type="text/css" media="all">




<link rel="preload" as="font"
      href="/themes/child_warehouse/assets/css/font-awesome/fonts/fontawesome-webfont.woff?v=4.7.0"
      type="font/woff" crossorigin="anonymous">
<link rel="preload" as="font"
      href="/themes/child_warehouse/assets/css/font-awesome/fonts/fontawesome-webfont.woff2?v=4.7.0"
      type="font/woff2" crossorigin="anonymous">


<link  rel="preload stylesheet"  as="style" href="/themes/child_warehouse/assets/css/font-awesome/css/font-awesome-preload.css"
       type="text/css" crossorigin="anonymous">





  

  <script>
        var PAYPLUG_DOMAIN = "https:\/\/secure.payplug.com";
        var elementorFrontendConfig = {"isEditMode":"","stretchedSectionContainer":"","instagramToken":false,"is_rtl":false,"ajax_csfr_token_url":"https:\/\/www.astronome.fr\/module\/iqitelementor\/Actions?process=handleCsfrToken&ajax=1"};
        var integratedPaymentError = "Paiement refus\u00e9, veuillez r\u00e9essayer.";
        var iqitTheme = {"rm_sticky":"0","rm_breakpoint":0,"op_preloader":"0","cart_style":"side","cart_confirmation":"modal","h_layout":"2","f_fixed":"","f_layout":"3","h_absolute":"0","h_sticky":"menu","hw_width":"inherit","mm_content":"accordion","hm_submenu_width":"default","h_search_type":"box","pl_lazyload":true,"pl_infinity":true,"pl_rollover":true,"pl_crsl_autoplay":false,"pl_slider_ld":5,"pl_slider_d":5,"pl_slider_t":4,"pl_slider_p":2,"pp_thumbs":"bottom","pp_zoom":"modalzoom","pp_image_layout":"carousel","pp_tabs":"tabha","pl_grid_qty":false};
        var iqitcountdown_days = "d.";
        var iqitmegamenu = {"sticky":"false","containerSelector":"#wrapper > .container"};
        var iqitwishlist = {"nbProducts":0};
        var is_sandbox_mode = false;
        var module_name = "payplug";
        var payplug_ajax_url = "https:\/\/www.astronome.fr\/module\/payplug\/ajax";
        var payplug_oney = true;
        var payplug_oney_loading_msg = "Chargement";
        var prestashop = {"cart":{"products":[],"totals":{"total":{"type":"total","label":"Total","amount":0,"value":"0,00\u00a0\u20ac"},"total_including_tax":{"type":"total","label":"Total TTC","amount":0,"value":"0,00\u00a0\u20ac"},"total_excluding_tax":{"type":"total","label":"Total HT :","amount":0,"value":"0,00\u00a0\u20ac"}},"subtotals":{"products":{"type":"products","label":"Sous-total","amount":0,"value":"0,00\u00a0\u20ac"},"discounts":null,"shipping":{"type":"shipping","label":"Livraison","amount":0,"value":""},"tax":null},"products_count":0,"summary_string":"0 articles","vouchers":{"allowed":0,"added":[]},"discounts":[],"minimalPurchase":0,"minimalPurchaseRequired":""},"currency":{"id":1,"name":"Euro","iso_code":"EUR","iso_code_num":"978","sign":"\u20ac"},"customer":{"lastname":null,"firstname":null,"email":null,"birthday":null,"newsletter":null,"newsletter_date_add":null,"optin":null,"website":null,"company":null,"siret":null,"ape":null,"is_logged":false,"gender":{"type":null,"name":null},"addresses":[]},"language":{"name":"Fran\u00e7ais (French)","iso_code":"fr","locale":"fr-FR","language_code":"fr","is_rtl":"0","date_format_lite":"d\/m\/Y","date_format_full":"d\/m\/Y H:i:s","id":1},"page":{"title":"","canonical":"https:\/\/www.astronome.fr\/lunettes-astronomiques\/3606-teleobjectif-redcat-91-wifd-william-optics.html","meta":{"title":"T\u00e9l\u00e9objectif RedCat 91 WIFD William Optics","description":"Le T\u00e9l\u00e9objectif RedCat 91 WIFD William Optics est le dernier n\u00e9 de la fabuleuse s\u00e9rie de t\u00e9l\u00e9objectifs RedCat d\u00e9di\u00e9s \u00e0 l'astrophotographie \u00e0 grand champ.","keywords":"","robots":"index"},"page_name":"product","body_classes":{"lang-fr":true,"lang-rtl":false,"country-FR":true,"currency-EUR":true,"layout-full-width":true,"page-product":true,"tax-display-enabled":true,"product-id-3606":true,"product-T\u00e9l\u00e9objectif RedCat 91 WIFD William Optics":true,"product-id-category-12":true,"product-id-manufacturer-5":true,"product-id-supplier-0":true,"product-available-for-order":true},"admin_notifications":[]},"shop":{"name":"L'Astronome","logo":"https:\/\/www.astronome.fr\/img\/logo-1685139979.jpg","stores_icon":"https:\/\/www.astronome.fr\/img\/logo_stores.png","favicon":"https:\/\/www.astronome.fr\/img\/favicon.ico"},"urls":{"base_url":"https:\/\/www.astronome.fr\/","current_url":"https:\/\/www.astronome.fr\/lunettes-astronomiques\/3606-teleobjectif-redcat-91-wifd-william-optics.html","shop_domain_url":"https:\/\/www.astronome.fr","img_ps_url":"https:\/\/www.astronome.fr\/img\/","img_cat_url":"https:\/\/www.astronome.fr\/img\/c\/","img_lang_url":"https:\/\/www.astronome.fr\/img\/l\/","img_prod_url":"https:\/\/www.astronome.fr\/img\/p\/","img_manu_url":"https:\/\/www.astronome.fr\/img\/m\/","img_sup_url":"https:\/\/www.astronome.fr\/img\/su\/","img_ship_url":"https:\/\/www.astronome.fr\/img\/s\/","img_store_url":"https:\/\/www.astronome.fr\/img\/st\/","img_col_url":"https:\/\/www.astronome.fr\/img\/co\/","img_url":"https:\/\/www.astronome.fr\/themes\/child_warehouse\/assets\/img\/","css_url":"https:\/\/www.astronome.fr\/themes\/child_warehouse\/assets\/css\/","js_url":"https:\/\/www.astronome.fr\/themes\/child_warehouse\/assets\/js\/","pic_url":"https:\/\/www.astronome.fr\/upload\/","pages":{"address":"https:\/\/www.astronome.fr\/adresse","addresses":"https:\/\/www.astronome.fr\/adresses","authentication":"https:\/\/www.astronome.fr\/connexion","cart":"https:\/\/www.astronome.fr\/panier","category":"https:\/\/www.astronome.fr\/index.php?controller=category","cms":"https:\/\/www.astronome.fr\/index.php?controller=cms","contact":"https:\/\/www.astronome.fr\/nous-contacter","discount":"https:\/\/www.astronome.fr\/reduction","guest_tracking":"https:\/\/www.astronome.fr\/suivi-commande-invite","history":"https:\/\/www.astronome.fr\/historique-commandes","identity":"https:\/\/www.astronome.fr\/identite","index":"https:\/\/www.astronome.fr\/","my_account":"https:\/\/www.astronome.fr\/mon-compte","order_confirmation":"https:\/\/www.astronome.fr\/confirmation-commande","order_detail":"https:\/\/www.astronome.fr\/index.php?controller=order-detail","order_follow":"https:\/\/www.astronome.fr\/suivi-commande","order":"https:\/\/www.astronome.fr\/commande","order_return":"https:\/\/www.astronome.fr\/index.php?controller=order-return","order_slip":"https:\/\/www.astronome.fr\/avoirs","pagenotfound":"https:\/\/www.astronome.fr\/page-introuvable","password":"https:\/\/www.astronome.fr\/recuperation-mot-de-passe","pdf_invoice":"https:\/\/www.astronome.fr\/index.php?controller=pdf-invoice","pdf_order_return":"https:\/\/www.astronome.fr\/index.php?controller=pdf-order-return","pdf_order_slip":"https:\/\/www.astronome.fr\/index.php?controller=pdf-order-slip","prices_drop":"https:\/\/www.astronome.fr\/promotions","product":"https:\/\/www.astronome.fr\/index.php?controller=product","search":"https:\/\/www.astronome.fr\/recherche","sitemap":"https:\/\/www.astronome.fr\/plan du site","stores":"https:\/\/www.astronome.fr\/magasins","supplier":"https:\/\/www.astronome.fr\/fournisseur","register":"https:\/\/www.astronome.fr\/connexion?create_account=1","order_login":"https:\/\/www.astronome.fr\/commande?login=1"},"alternative_langs":[],"theme_assets":"\/themes\/child_warehouse\/assets\/","actions":{"logout":"https:\/\/www.astronome.fr\/?mylogout="},"no_picture_image":{"bySize":{"small_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-small_default.jpg","width":98,"height":127},"cart_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-cart_default.jpg","width":125,"height":162},"home_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-home_default.jpg","width":236,"height":305},"large_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-large_default.jpg","width":381,"height":492},"medium_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-medium_default.jpg","width":452,"height":584},"thickbox_default":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-thickbox_default.jpg","width":1100,"height":1422}},"small":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-small_default.jpg","width":98,"height":127},"medium":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-large_default.jpg","width":381,"height":492},"large":{"url":"https:\/\/www.astronome.fr\/img\/p\/fr-default-thickbox_default.jpg","width":1100,"height":1422},"legend":""}},"configuration":{"display_taxes_label":true,"display_prices_tax_incl":true,"is_catalog":false,"show_prices":true,"opt_in":{"partner":false},"quantity_discount":{"type":"discount","label":"Remise sur prix unitaire"},"voucher_enabled":0,"return_enabled":0},"field_required":[],"breadcrumb":{"links":[{"title":"Accueil","url":"https:\/\/www.astronome.fr\/"},{"title":"Lunettes astronomiques","url":"https:\/\/www.astronome.fr\/12-lunettes-astronomiques"},{"title":"T\u00e9l\u00e9objectif RedCat 91 WIFD William Optics","url":"https:\/\/www.astronome.fr\/lunettes-astronomiques\/3606-teleobjectif-redcat-91-wifd-william-optics.html"}],"count":3},"link":{"protocol_link":"https:\/\/","protocol_content":"https:\/\/"},"time":1731918112,"static_token":"0cf0ed0367fc06840609b7167cf23032","token":"427da2ad027bcc71dfd6810553ea4307","debug":false};
        var psemailsubscription_subscription = "https:\/\/www.astronome.fr\/module\/ps_emailsubscription\/subscription";
      </script>



  
<script type="text/javascript">
    var lgcookieslaw_consent_mode = 0;
    var lgcookieslaw_banner_url_ajax_controller = "https://www.astronome.fr/module/lgcookieslaw/ajax";     var lgcookieslaw_cookie_values = {"lgcookieslaw_purpose_1":true,"lgcookieslaw_purpose_2":true,"lgcookieslaw_purpose_3":true,"lgcookieslaw_purpose_4":true,"lgcookieslaw_purpose_5":true,"lgcookieslaw_user_consent_consent_date":"2024-10-29 17:38:55","lgcookieslaw_user_consent_download_hash":"a02986edfd961b7d230d2fcab4be16bd","lgcookieslaw_accepted_purposes":[1,2,3,4,5],"lgcookieslaw_user_consent_download_url":"https:\/\/www.astronome.fr\/module\/lgcookieslaw\/download?id_shop=1&download_hash=a02986edfd961b7d230d2fcab4be16bd"};     var lgcookieslaw_saved_preferences = 1;
    var lgcookieslaw_ajax_calls_token = "f9d19a395b7b65ea3ee113440a05a380";
    var lgcookieslaw_reload = 1;
    var lgcookieslaw_block_navigation = 0;
    var lgcookieslaw_banner_position = 2;
    var lgcookieslaw_show_fixed_button = 0;
    var lgcookieslaw_save_user_consent = 1;
    var lgcookieslaw_reject_cookies_when_closing_banner = 0;
</script>


<script async src="https://www.googletagmanager.com/gtag/js?id=G-0HFST6HJ09"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag(
    'config',
    'G-0HFST6HJ09',
    {
      'debug_mode':false
      , 'anonymize_ip': true                }
  );
</script>




    
            <meta property="product:pretax_price:amount" content="2749.166667">
        <meta property="product:pretax_price:currency" content="EUR">
        <meta property="product:price:amount" content="3299">
        <meta property="product:price:currency" content="EUR">
                <meta property="product:weight:value" content="25.000000">
        <meta property="product:weight:units" content="kg">
    
    

    </head>
</html>
	`
	parser := Parser{&ParserConfig{FilteredFileExtensions: []string{".css"}}}
	urls := []string{}
	r := io.NopCloser(strings.NewReader(raw_html)) // r type is io.ReadCloser	if err != nil {
	err := parser.GetAllUrls(r, &urls)
	if err != nil {
		t.Fail()
	}
	t.Log(urls)
	index := slices.IndexFunc(urls, func(s string) bool {
		return strings.Contains("https://www.astronome.fr/themes/child_warehouse/assets/cache/theme-f7d352620.css", s)
	})
	if index == -1 {
		t.Fatalf("not good")
	}

}
