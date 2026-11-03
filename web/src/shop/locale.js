import { ref } from 'vue'

const KEY = 'blink.shop.locale'

const stored = localStorage.getItem(KEY)
export const shopLocale = ref(stored === 'ar' ? 'ar' : 'en')

export function setShopLocale(locale) {
  shopLocale.value = locale
  localStorage.setItem(KEY, locale)
}

const ui = {
  en: {
    eta: '10–15 min',
    searchPlaceholder: 'Search bananas, milk, snacks',
    searchStore: 'Search the store',
    searchHint: 'Try milk, avocado, or pizza.',
    noMatch: 'Nothing matches “{query}”.',
    shop: 'Shop',
    search: 'Search',
    basket: 'Basket',
    you: 'You',
    viewBasket: 'View basket',
    tagline: 'Your neighborhood store, on your doorstep in 10–15 minutes.',
    legal: 'Legal',
    rights: '© 2026 Blink. All Rights Reserved.',
    deliveringTo: 'Delivering to',
    fromNearby: 'From a store nearby',
    hero: 'Groceries at your door in minutes.',
    freeFrom: 'Free delivery from {amount}.',
    seeAll: 'See all',
    somethingToDrink: 'Something to drink',
    addressTitle: 'Delivery address',
    addressHelp: 'We’ll drop the order at this door. Saved on this device only.',
    addressPlaceholder: 'Street and neighbourhood',
    saveAddress: 'Save address',
    saved: 'Saved',
    deal: 'Deal',
    addToBasket: 'Add to basket',
    inBasket: '{n} in basket',
    oneProduct: '1 product',
    nProducts: '{n} products',
    goesWith: 'Goes with it',
    emptyTitle: 'Your basket is empty',
    emptyBody: 'The store is open. 10–15 minutes to your door.',
    startShopping: 'Start shopping',
    freeUnlocked: 'Free delivery unlocked.',
    addForFree: 'Add {amount} for free delivery.',
    subtotal: 'Subtotal',
    delivery: 'Delivery',
    free: 'Free',
    total: 'Total',
    checkout: 'Checkout',
    checkoutCta: 'Checkout · {amount}',
    oneItem: '1 item',
    nItems: '{n} items',
    arriving: 'arriving in 10–15 min',
    name: 'Name',
    phone: 'Phone',
    address: 'Address',
    window: 'Delivery window',
    windowAsap: 'As soon as possible · 10–15 min',
    window30: 'In about 30 min',
    window60: 'In about an hour',
    note: 'Note for the courier',
    notePlaceholder: 'Door code, floor, leave by the mat',
    placing: 'Placing order…',
    placeOrder: 'Place order · {amount}',
    preview: 'Preview only. Nothing is charged and no courier is dispatched.',
    emptyError: 'Your basket is empty.',
    guest: 'Shopping as a guest on this device.',
    orders: 'Orders',
    noOrders: 'No orders yet. They’ll show up here after checkout.',
    orderItems: '{n} items',
    orderLabel: 'Order {id}',
    onTheWay: 'On the way',
    orderLine: '{window} to {address}.',
    totalAmount: 'Total {amount}',
    savedDevice: 'Saved on this device. No payment was taken.',
    backToStore: 'Back to the store',
    signIn: 'Sign in',
    createAccount: 'Create account',
    email: 'Email',
    password: 'Password',
    signOut: 'Sign out',
    hello: 'Hello, {name}',
    yourDetails: 'Your details',
    haveAccount: 'Already have an account? Sign in',
    needAccount: 'New here? Create an account',
    badLogin: 'That email and password do not match.',
    emailTaken: 'An account with that email already exists.',
    signingIn: 'Signing in…',
    creating: 'Creating account…',
    profileNote: 'This account stays on this device.',
    continueGoogle: 'Continue with Google',
    continueFacebook: 'Continue with Facebook',
    continueApple: 'Continue with Apple',
    or: 'or',
    signedInWith: 'Signed in with {provider}',
  },
  ar: {
    eta: '١٠–١٥ د',
    searchPlaceholder: 'ابحث عن موز أو حليب أو سناك',
    searchStore: 'ابحث في المتجر',
    searchHint: 'جرّب حليب أو أفوكادو أو بيتزا.',
    noMatch: 'لا نتائج لـ «{query}».',
    shop: 'المتجر',
    search: 'بحث',
    basket: 'السلة',
    you: 'حسابي',
    viewBasket: 'عرض السلة',
    tagline: 'متجرك في الحي، على عتبة بابك خلال ١٠–١٥ دقيقة.',
    legal: 'قانوني',
    rights: '© 2026 Blink. جميع الحقوق محفوظة.',
    deliveringTo: 'التوصيل إلى',
    fromNearby: 'من متجر قريب',
    hero: 'مقاضي إلى بابك خلال دقائق.',
    freeFrom: 'توصيل مجاني من {amount}.',
    seeAll: 'عرض الكل',
    somethingToDrink: 'شيء للشرب',
    addressTitle: 'عنوان التوصيل',
    addressHelp: 'نوصل الطلب إلى هذا الباب. يُحفظ على هذا الجهاز فقط.',
    addressPlaceholder: 'الشارع والحي',
    saveAddress: 'حفظ العنوان',
    saved: 'تم الحفظ',
    deal: 'عرض',
    addToBasket: 'أضف إلى السلة',
    inBasket: '{n} في السلة',
    oneProduct: 'منتج واحد',
    nProducts: '{n} منتجات',
    goesWith: 'يناسبها',
    emptyTitle: 'سلتك فارغة',
    emptyBody: 'المتجر مفتوح. ١٠–١٥ دقيقة إلى بابك.',
    startShopping: 'ابدأ التسوق',
    freeUnlocked: 'التوصيل المجاني مفتوح.',
    addForFree: 'أضف {amount} للتوصيل المجاني.',
    subtotal: 'المجموع',
    delivery: 'التوصيل',
    free: 'مجاني',
    total: 'الإجمالي',
    checkout: 'الدفع',
    checkoutCta: 'الدفع · {amount}',
    oneItem: 'صنف واحد',
    nItems: '{n} أصناف',
    arriving: 'يصل خلال ١٠–١٥ د',
    name: 'الاسم',
    phone: 'الهاتف',
    address: 'العنوان',
    window: 'وقت التوصيل',
    windowAsap: 'في أقرب وقت · ١٠–١٥ د',
    window30: 'بعد نحو ٣٠ دقيقة',
    window60: 'بعد نحو ساعة',
    note: 'ملاحظة للمندوب',
    notePlaceholder: 'رمز الباب، الطابق، أو اتركه عند العتبة',
    placing: 'جارٍ إرسال الطلب…',
    placeOrder: 'أرسل الطلب · {amount}',
    preview: 'معاينة فقط. لا يُخصم شيء ولا يُرسل مندوب.',
    emptyError: 'سلتك فارغة.',
    guest: 'تتسوق كزائر على هذا الجهاز.',
    orders: 'الطلبات',
    noOrders: 'لا طلبات بعد. تظهر هنا بعد الدفع.',
    orderItems: '{n} أصناف',
    orderLabel: 'طلب {id}',
    onTheWay: 'في الطريق',
    orderLine: '{window} إلى {address}.',
    totalAmount: 'الإجمالي {amount}',
    savedDevice: 'محفوظ على هذا الجهاز. لم يُؤخذ أي مبلغ.',
    backToStore: 'العودة إلى المتجر',
    signIn: 'دخول',
    createAccount: 'إنشاء حساب',
    email: 'البريد',
    password: 'كلمة المرور',
    signOut: 'خروج',
    hello: 'مرحبًا {name}',
    yourDetails: 'بياناتك',
    haveAccount: 'لديك حساب؟ سجّل الدخول',
    needAccount: 'جديد هنا؟ أنشئ حسابًا',
    badLogin: 'البريد أو كلمة المرور غير مطابقة.',
    emailTaken: 'يوجد حساب بهذا البريد.',
    signingIn: 'جارٍ الدخول…',
    creating: 'جارٍ إنشاء الحساب…',
    profileNote: 'هذا الحساب يبقى على هذا الجهاز.',
    continueGoogle: 'المتابعة مع Google',
    continueFacebook: 'المتابعة مع Facebook',
    continueApple: 'المتابعة مع Apple',
    or: 'أو',
    signedInWith: 'دخول عبر {provider}',
  },
}

const categoriesAr = {
  deals: 'العروض',
  fruit: 'فواكه وخضار',
  'fresh-fruit': 'فواكه طازجة',
  vegetables: 'خضار',
  dairy: 'ألبان وبيض',
  'milk-yogurt': 'حليب ولبن',
  'cheese-eggs': 'جبن وبيض',
  bakery: 'مخبوزات',
  bread: 'خبز',
  pastries: 'معجنات',
  meat: 'لحوم وأسماك',
  'fresh-meat': 'لحوم',
  fish: 'أسماك',
  drinks: 'مشروبات',
  'soft-drinks': 'مشروبات باردة',
  'coffee-beer': 'قهوة وبيرة',
  snacks: 'سناك',
  savoury: 'مالح',
  sweet: 'حلو',
  ready: 'جاهز للأكل',
  meals: 'وجبات',
  pizza: 'بيتزا',
  frozen: 'مجمدات',
  'frozen-meals': 'وجبات مجمدة',
  'ice-cream': 'آيس كريم',
  home: 'منزل',
  paper: 'ورق',
  cleaning: 'تنظيف',
}

const productsAr = {
  bananas: {
    name: 'موز',
    pack: 'حوالي 1 كغ',
    unitPrice: '1,49 € / كغ',
    blurb: 'موز أصفر ناضج، للفطور أو كسناك.',
  },
  avocados: {
    name: 'أفوكادو هاس',
    pack: 'قطعتان',
    unitPrice: '1,25 € / حبة',
    blurb: 'كريمي وناضج بما يكفي للهرس الليلة.',
  },
  strawberries: {
    name: 'فراولة',
    pack: '250 غ',
    unitPrice: '11,96 € / كغ',
    blurb: 'فراولة حلوة، أفضل ما تكون يوم وصولها.',
  },
  tomatoes: {
    name: 'طماطم كرزية',
    pack: '250 غ',
    unitPrice: '8,76 € / كغ',
    blurb: 'طماطم صغيرة بطعم حلو وواضح.',
  },
  spinach: {
    name: 'سبانخ صغيرة',
    pack: '100 غ',
    unitPrice: '17,90 € / كغ',
    blurb: 'أوراق طرية للسلطة أو البيض أو سوتيه سريع.',
  },
  blueberries: {
    name: 'توت أزرق',
    pack: '125 غ',
    unitPrice: '19,92 € / كغ',
    blurb: 'حبات ممتلئة للبن أو الشوفان أو من العلبة.',
  },
  milk: {
    name: 'حليب كامل الدسم',
    pack: '1 ل',
    unitPrice: '1,29 € / ل',
    blurb: 'حليب طازج للقهوة والحبوب والطبخ.',
  },
  yogurt: {
    name: 'لبن يوناني',
    pack: '500 غ',
    unitPrice: '3,78 € / كغ',
    blurb: 'لبن كثيف سادة. دسم 10٪.',
  },
  eggs: {
    name: 'بيض حر',
    pack: '10 حبات',
    unitPrice: '0,33 € / حبة',
    blurb: 'بيض بني كبير من دجاج طليق.',
  },
  butter: {
    name: 'زبدة مملحة',
    pack: '250 غ',
    unitPrice: '9,96 € / كغ',
    blurb: 'زبدة مخمرة للخبز والصلصات والخبز.',
  },
  cheddar: {
    name: 'شيدر معتّق',
    pack: '200 غ',
    unitPrice: '13,95 € / كغ',
    blurb: 'معتّق 12 شهرًا. حاد ومتفتت، لذيذ على الخبز.',
  },
  sourdough: {
    name: 'خبز عجينة مخمرة',
    pack: '500 غ',
    unitPrice: '5,78 € / كغ',
    blurb: 'خُبز هذا الصباح. قشرة مقرمشة ولب مفتوح.',
  },
  croissants: {
    name: 'كرواسون بالزبدة',
    pack: 'قطعتان',
    unitPrice: '1,00 € / حبة',
    blurb: 'بالزبدة، هش، وما زال دافئًا قليلًا.',
  },
  bagels: {
    name: 'بيغل بالسمسم',
    pack: '4 حبات',
    unitPrice: '0,62 € / حبة',
    blurb: 'بيغل مطاطي. حمّصه وأضف الشيدر.',
  },
  chicken: {
    name: 'صدر دجاج',
    pack: '400 غ',
    unitPrice: '12,48 € / كغ',
    blurb: 'شرائح بدون جلد، عُبّئت اليوم.',
  },
  salmon: {
    name: 'فيليه سلمون',
    pack: '200 غ',
    unitPrice: '27,45 € / كغ',
    blurb: 'سلمون أطلسي منزوع الشوك، جاهز للمقلاة.',
  },
  mince: {
    name: 'لحم بقر مفروم',
    pack: '500 غ',
    unitPrice: '11,98 € / كغ',
    blurb: 'دهن 12٪. يكفي بولونيز للعائلة.',
  },
  water: {
    name: 'مياه غازية',
    pack: '6 × 330 مل',
    unitPrice: '1,26 € / ل',
    blurb: 'مياه معدنية منعشة بفقاعات ناعمة.',
  },
  juice: {
    name: 'عصير برتقال',
    pack: '1 ل',
    unitPrice: '2,19 € / ل',
    blurb: 'ليس من مركز. رُجّه أولًا.',
  },
  oat: {
    name: 'مشروب شوفان',
    pack: '1 ل',
    unitPrice: '1,89 € / ل',
    blurb: 'شوفان باريستا يُبخّر دون أن ينفصل.',
  },
  coffee: {
    name: 'قهوة باردة',
    pack: '250 مل',
    unitPrice: '9,96 € / ل',
    blurb: 'بلا سكر، نُقعت ليلًا وتُشرب باردة.',
  },
  ipa: {
    name: 'بيرة خفيفة',
    pack: '4 × 330 مل',
    unitPrice: '4,54 € / ل',
    blurb: 'حمضية وخفيفة. 4.2٪.',
  },
  crisps: {
    name: 'رقائق مملحة',
    pack: '150 غ',
    unitPrice: '11,27 € / كغ',
    blurb: 'بطاطس سميكة مقلية ومملحة. ولا شيء آخر.',
  },
  chocolate: {
    name: 'شوكولاتة داكنة',
    pack: '100 غ',
    unitPrice: '18,90 € / كغ',
    blurb: 'كاكاو 70٪. قضم نظيف، ليست حلوة جدًا.',
  },
  nuts: {
    name: 'مكسرات محمصة',
    pack: '200 غ',
    unitPrice: '17,45 € / كغ',
    blurb: 'لوز وكاجو وبندق، مملحة قليلًا.',
  },
  'pasta-meal': {
    name: 'باستا دجاج بالبيستو',
    pack: '400 غ',
    unitPrice: '13,73 € / كغ',
    blurb: 'تُسخّن لأربع دقائق. العشاء جاهز.',
  },
  pizza: {
    name: 'بيتزا مارغريتا',
    pack: '430 غ',
    unitPrice: '11,60 € / كغ',
    blurb: 'قاعدة بنكهة الفرن، موتزاريلا وريحان.',
  },
  dumplings: {
    name: 'زلابية بازلاء ونعناع',
    pack: '400 غ',
    unitPrice: '9,98 € / كغ',
    blurb: 'من المجمد إلى المقلاة في ثماني دقائق.',
  },
  icecream: {
    name: 'آيس كريم فانيليا',
    pack: '460 مل',
    unitPrice: '7,59 € / ل',
    blurb: 'فانيليا مدغشقر مخفوقة بقشطة حقيقية.',
  },
  towel: {
    name: 'مناديل مطبخ',
    pack: '4 لفات',
    unitPrice: '0,57 € / لفة',
    blurb: 'طبقتان قويتان للانسكابات وكل شيء آخر.',
  },
  soap: {
    name: 'سائل جلي',
    pack: '500 مل',
    unitPrice: '3,18 € / ل',
    blurb: 'رائحة ليمون. يزيل الدهن بهدوء.',
  },
}

const pagesAr = {
  about: {
    slug: 'about',
    title: 'عن بلينك',
    lead: 'مقاضي من متجر قريب، إلى بابك خلال دقائق.',
    sections: [
      {
        heading: 'متجر، لا مستودع',
        body: 'بلينك يأخذ رف متجر قريب منك ويحضره خلال ١٠–١٥ دقيقة. فواكه وحليب وعشاء ولفة المطبخ التي نسيتها.',
      },
      {
        heading: 'هذه المعاينة',
        body: 'الكتالوج والسلة والطلبات تبقى على جهازك. لا يُخصم شيء ولا يُرسل مندوب. الصفحة هنا لتجرّب المتجر قبل أن يفتح فعليًا.',
      },
    ],
  },
  delivery: {
    slug: 'delivery',
    title: 'التوصيل',
    lead: 'اختر وقتًا عند الدفع. نقصد الباب، لا المستودع.',
    sections: [
      {
        heading: 'كم يأخذ',
        body: 'في أقرب وقت يعني ١٠–١٥ دقيقة. يمكنك أيضًا طلب نحو ٣٠ دقيقة أو نحو ساعة إن لم تكن في البيت بعد.',
      },
      {
        heading: 'كم يكلف',
        body: 'التوصيل 1,90 €. يصبح مجانيًا عندما تصل السلة إلى 29,00 €. الشريط في السلة يريك المسافة المتبقية.',
      },
      {
        heading: 'إلى أين',
        body: 'العنوان الذي تحفظه هو الباب الذي نستخدمه. أضف ملاحظة للرمز أو الطابق أو «اتركه عند العتبة».',
      },
    ],
  },
  help: {
    slug: 'help',
    title: 'مساعدة',
    lead: 'النسخة القصيرة لطريقة عمل هذا المتجر.',
    sections: [
      {
        heading: 'كيف تجد شيئًا',
        body: 'افتح قسمًا ثم قسمًا فرعيًا، أو استخدم البحث. العروض هي المنتجات ذات العلامة الوردية والسعر المشطوب.',
      },
      {
        heading: 'سلتك',
        body: 'علامة الزائد على المنتج تضيف واحدًا. افتح السلة لتغيير الكمية ورؤية التوصيل والدفع. السلة تبقى على هذا الجهاز.',
      },
      {
        heading: 'الطلبات',
        body: 'بعد الدفع يظهر الطلب تحت حسابي. إنه إيصال معاينة. لا يُؤخذ أي مبلغ.',
      },
      {
        heading: 'سؤال حقيقي',
        body: 'اكتب إلى hello@clivern.com وسنرد عليك.',
      },
    ],
  },
  terms: {
    slug: 'terms',
    title: 'الشروط',
    lead: 'قواعد معاينة متجر بلينك.',
    sections: [
      {
        heading: 'معاينة، لا بيع',
        body: 'إرسال طلب هنا لا يشتري شيئًا. لا يُؤخذ مبلغ، ولا يُحجز متجر، ولا يُرسل مندوب.',
      },
      {
        heading: 'الكتالوج',
        body: 'الأسعار والعبوات والمنتجات بيانات تجريبية. قد تتغير، وليست عرضًا من متجر.',
      },
      {
        heading: 'استخدامك',
        body: 'استخدم المتجر كما يظهر. لا تعتمد على طلب في هذه المعاينة كإيصال أو حجز أو إثبات توصيل.',
      },
    ],
  },
  privacy: {
    slug: 'privacy',
    title: 'الخصوصية',
    lead: 'ما يحتفظ به هذا المتجر، وأين يحتفظ به.',
    sections: [
      {
        heading: 'على هذا الجهاز',
        body: 'عنوان التوصيل والسلة وطلبات المعاينة السابقة تُخزَّن في هذا المتصفح. لا تُرسل إلى حساب بلينك، لأن هذا المتجر لا يطلب تسجيل الدخول.',
      },
      {
        heading: 'ما تكتبه عند الدفع',
        body: 'الاسم والهاتف والعنوان وملاحظة المندوب تُحفظ مع الطلب على هذا الجهاز ليظهر الإيصال. مسح بيانات الموقع يزيلها.',
      },
      {
        heading: 'ملفات تعريف الارتباط',
        body: 'معاينة المتجر لا تضيف تتبّعًا خاصًا بها. بقية بلينك تطلب خيارات الكوكيز بشكل منفصل.',
      },
    ],
  },
}

export function t(key, vars = {}) {
  const table = ui[shopLocale.value]
  let text = table[key] ?? ui.en[key]
  for (const [name, value] of Object.entries(vars)) {
    text = text.replaceAll(`{${name}}`, String(value))
  }
  return text
}

export function categoryLabel(category) {
  if (shopLocale.value === 'ar') return categoriesAr[category.slug]
  return category.name
}

export function localizeProduct(product) {
  if (shopLocale.value !== 'ar') return product
  return { ...product, ...productsAr[product.sku] }
}

export function localizePage(page) {
  if (shopLocale.value !== 'ar') return page
  return pagesAr[page.slug]
}

export function productsLabel(count) {
  if (count === 1) return t('oneProduct')
  return t('nProducts', { n: count })
}
