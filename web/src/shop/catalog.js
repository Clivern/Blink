import { reactive } from 'vue'

const CATALOG_KEY = 'blink.shop.catalog'

export const DELIVERY_CENTS = 190
export const FREE_DELIVERY_OVER = 2900

const seedCategories = [
  { slug: 'deals', name: 'Deals', emoji: '⚡', tint: '#ffe1ee', children: [] },
  {
    slug: 'fruit',
    name: 'Fruit & veg',
    emoji: '🥑',
    tint: '#e4f6d8',
    children: [
      { slug: 'fresh-fruit', name: 'Fresh fruit', emoji: '🍌', tint: '#fff4c8' },
      { slug: 'vegetables', name: 'Vegetables', emoji: '🥬', tint: '#e3f6df' },
    ],
  },
  {
    slug: 'dairy',
    name: 'Dairy & eggs',
    emoji: '🥛',
    tint: '#e7f2ff',
    children: [
      { slug: 'milk-yogurt', name: 'Milk & yogurt', emoji: '🥛', tint: '#eef6ff' },
      { slug: 'cheese-eggs', name: 'Cheese & eggs', emoji: '🧀', tint: '#fff0c8' },
    ],
  },
  {
    slug: 'bakery',
    name: 'Bakery',
    emoji: '🥐',
    tint: '#fff0d6',
    children: [
      { slug: 'bread', name: 'Bread', emoji: '🍞', tint: '#f8ead2' },
      { slug: 'pastries', name: 'Pastries', emoji: '🥐', tint: '#fff1d4' },
    ],
  },
  {
    slug: 'meat',
    name: 'Meat & fish',
    emoji: '🐟',
    tint: '#ffe8e4',
    children: [
      { slug: 'fresh-meat', name: 'Meat', emoji: '🥩', tint: '#ffe3e0' },
      { slug: 'fish', name: 'Fish', emoji: '🍣', tint: '#ffe4ea' },
    ],
  },
  {
    slug: 'drinks',
    name: 'Drinks',
    emoji: '🥤',
    tint: '#e8f7f4',
    children: [
      { slug: 'soft-drinks', name: 'Soft drinks', emoji: '🍊', tint: '#fff1d8' },
      { slug: 'coffee-beer', name: 'Coffee & beer', emoji: '☕', tint: '#f0e6da' },
    ],
  },
  {
    slug: 'snacks',
    name: 'Snacks',
    emoji: '🍫',
    tint: '#f6e8ff',
    children: [
      { slug: 'savoury', name: 'Savoury', emoji: '🥔', tint: '#fff6d4' },
      { slug: 'sweet', name: 'Sweet', emoji: '🍫', tint: '#f6e4d8' },
    ],
  },
  {
    slug: 'ready',
    name: 'Ready to eat',
    emoji: '🍝',
    tint: '#ffe9d6',
    children: [
      { slug: 'meals', name: 'Meals', emoji: '🍝', tint: '#eef8e4' },
      { slug: 'pizza', name: 'Pizza', emoji: '🍕', tint: '#fff0dc' },
    ],
  },
  {
    slug: 'frozen',
    name: 'Frozen',
    emoji: '🍨',
    tint: '#e8f0ff',
    children: [
      { slug: 'frozen-meals', name: 'Frozen meals', emoji: '🥟', tint: '#e7f8ef' },
      { slug: 'ice-cream', name: 'Ice cream', emoji: '🍨', tint: '#f7f1e6' },
    ],
  },
  {
    slug: 'home',
    name: 'Household',
    emoji: '🧻',
    tint: '#f3f3f1',
    children: [
      { slug: 'paper', name: 'Paper', emoji: '🧻', tint: '#f4f4f2' },
      { slug: 'cleaning', name: 'Cleaning', emoji: '🧴', tint: '#e5f6f2' },
    ],
  },
]

const seedProducts = [
  {
    sku: 'bananas',
    name: 'Bananas',
    brand: 'Field',
    pack: 'about 1 kg',
    cents: 149,
    category: 'fresh-fruit',
    emoji: '🍌',
    tint: '#fff4c8',
    unitPrice: '1,49 € / kg',
    blurb: 'Ripe yellow bananas, ready for breakfast or a snack.',
    tags: ['popular'],
  },
  {
    sku: 'avocados',
    name: 'Hass avocados',
    brand: 'Field',
    pack: '2 pcs',
    cents: 249,
    wasCents: 329,
    category: 'fresh-fruit',
    emoji: '🥑',
    tint: '#e5f6d4',
    unitPrice: '1,25 € / pc',
    blurb: 'Creamy and just ripe enough to smash tonight.',
    tags: ['popular'],
  },
  {
    sku: 'strawberries',
    name: 'Strawberries',
    brand: 'Field',
    pack: '250 g',
    cents: 299,
    category: 'fresh-fruit',
    emoji: '🍓',
    tint: '#ffe4ea',
    unitPrice: '11,96 € / kg',
    blurb: 'Sweet berries, best eaten the day they arrive.',
    tags: [],
  },
  {
    sku: 'tomatoes',
    name: 'Cherry tomatoes',
    brand: 'Field',
    pack: '250 g',
    cents: 219,
    category: 'vegetables',
    emoji: '🍅',
    tint: '#ffe8e0',
    unitPrice: '8,76 € / kg',
    blurb: 'Snack-size tomatoes with a bright, sweet bite.',
    tags: ['popular'],
  },
  {
    sku: 'spinach',
    name: 'Baby spinach',
    brand: 'Field',
    pack: '100 g',
    cents: 179,
    category: 'vegetables',
    emoji: '🥬',
    tint: '#e3f6df',
    unitPrice: '17,90 € / kg',
    blurb: 'Tender leaves for salads, eggs, or a quick wilt.',
    tags: [],
  },
  {
    sku: 'blueberries',
    name: 'Blueberries',
    brand: 'Field',
    pack: '125 g',
    cents: 249,
    wasCents: 299,
    category: 'fresh-fruit',
    emoji: '🫐',
    tint: '#ece7ff',
    unitPrice: '19,92 € / kg',
    blurb: 'Plump berries for yogurt, porridge, or the box.',
    tags: [],
  },
  {
    sku: 'milk',
    name: 'Whole milk',
    brand: 'Daybreak',
    pack: '1 L',
    cents: 129,
    category: 'milk-yogurt',
    emoji: '🥛',
    tint: '#eef6ff',
    unitPrice: '1,29 € / L',
    blurb: 'Fresh whole milk for coffee, cereal, and cooking.',
    tags: ['popular'],
  },
  {
    sku: 'yogurt',
    name: 'Greek yogurt',
    brand: 'Daybreak',
    pack: '500 g',
    cents: 189,
    wasCents: 229,
    category: 'milk-yogurt',
    emoji: '🥣',
    tint: '#f4f7ff',
    unitPrice: '3,78 € / kg',
    blurb: 'Thick, plain yogurt. 10% fat.',
    tags: ['popular'],
  },
  {
    sku: 'eggs',
    name: 'Free-range eggs',
    brand: 'Daybreak',
    pack: '10 pcs',
    cents: 329,
    category: 'cheese-eggs',
    emoji: '🥚',
    tint: '#fff6e4',
    unitPrice: '0,33 € / pc',
    blurb: 'Large brown eggs from free-range hens.',
    tags: ['popular'],
  },
  {
    sku: 'butter',
    name: 'Salted butter',
    brand: 'Daybreak',
    pack: '250 g',
    cents: 249,
    category: 'cheese-eggs',
    emoji: '🧈',
    tint: '#fff3c9',
    unitPrice: '9,96 € / kg',
    blurb: 'Cultured butter for toast, sauces, and baking.',
    tags: [],
  },
  {
    sku: 'cheddar',
    name: 'Mature cheddar',
    brand: 'Daybreak',
    pack: '200 g',
    cents: 279,
    category: 'cheese-eggs',
    emoji: '🧀',
    tint: '#fff0c8',
    unitPrice: '13,95 € / kg',
    blurb: 'Aged 12 months. Sharp, crumbly, good on toast.',
    tags: [],
  },
  {
    sku: 'sourdough',
    name: 'Sourdough loaf',
    brand: 'Oven',
    pack: '500 g',
    cents: 289,
    category: 'bread',
    emoji: '🍞',
    tint: '#f8ead2',
    unitPrice: '5,78 € / kg',
    blurb: 'Baked this morning. Crackly crust, open crumb.',
    tags: ['popular'],
  },
  {
    sku: 'croissants',
    name: 'Butter croissants',
    brand: 'Oven',
    pack: '2 pcs',
    cents: 199,
    wasCents: 249,
    category: 'pastries',
    emoji: '🥐',
    tint: '#fff1d4',
    unitPrice: '1,00 € / pc',
    blurb: 'All-butter, flaky, and still a little warm.',
    tags: [],
  },
  {
    sku: 'bagels',
    name: 'Sesame bagels',
    brand: 'Oven',
    pack: '4 pcs',
    cents: 249,
    category: 'bread',
    emoji: '🥯',
    tint: '#f6e7cf',
    unitPrice: '0,62 € / pc',
    blurb: 'Chewy bagels. Toast them and add the cheddar.',
    tags: [],
  },
  {
    sku: 'chicken',
    name: 'Chicken breast',
    brand: 'Harbor',
    pack: '400 g',
    cents: 499,
    category: 'fresh-meat',
    emoji: '🍗',
    tint: '#ffe8df',
    unitPrice: '12,48 € / kg',
    blurb: 'Skinless fillets, packed today.',
    tags: ['popular'],
  },
  {
    sku: 'salmon',
    name: 'Salmon fillet',
    brand: 'Harbor',
    pack: '200 g',
    cents: 549,
    wasCents: 649,
    category: 'fish',
    emoji: '🍣',
    tint: '#ffe4ea',
    unitPrice: '27,45 € / kg',
    blurb: 'Atlantic salmon, pin-boned, ready for the pan.',
    tags: [],
  },
  {
    sku: 'mince',
    name: 'Beef mince',
    brand: 'Harbor',
    pack: '500 g',
    cents: 599,
    category: 'fresh-meat',
    emoji: '🥩',
    tint: '#ffe3e0',
    unitPrice: '11,98 € / kg',
    blurb: '12% fat. Enough for a family bolognese.',
    tags: [],
  },
  {
    sku: 'water',
    name: 'Sparkling water',
    brand: 'North',
    pack: '6 × 330 ml',
    cents: 249,
    wasCents: 319,
    category: 'soft-drinks',
    emoji: '🫧',
    tint: '#e7f6ff',
    unitPrice: '1,26 € / L',
    blurb: 'Crisp mineral water with fine bubbles.',
    tags: ['popular'],
  },
  {
    sku: 'juice',
    name: 'Orange juice',
    brand: 'North',
    pack: '1 L',
    cents: 219,
    category: 'soft-drinks',
    emoji: '🍊',
    tint: '#fff1d8',
    unitPrice: '2,19 € / L',
    blurb: 'Not from concentrate. Shake it first.',
    tags: [],
  },
  {
    sku: 'oat',
    name: 'Oat drink',
    brand: 'North',
    pack: '1 L',
    cents: 189,
    category: 'soft-drinks',
    emoji: '🌾',
    tint: '#f4efe4',
    unitPrice: '1,89 € / L',
    blurb: 'Barista oat drink that steams without splitting.',
    tags: ['popular'],
  },
  {
    sku: 'coffee',
    name: 'Cold brew',
    brand: 'North',
    pack: '250 ml',
    cents: 249,
    category: 'coffee-beer',
    emoji: '☕',
    tint: '#f0e6da',
    unitPrice: '9,96 € / L',
    blurb: 'Unsweetened, steeped overnight, served cold.',
    tags: [],
  },
  {
    sku: 'ipa',
    name: 'Session IPA',
    brand: 'North',
    pack: '4 × 330 ml',
    cents: 599,
    category: 'coffee-beer',
    emoji: '🍺',
    tint: '#fff4d2',
    unitPrice: '4,54 € / L',
    blurb: 'Citrusy and light. 4.2% ABV.',
    tags: [],
  },
  {
    sku: 'crisps',
    name: 'Salted crisps',
    brand: 'Blink',
    pack: '150 g',
    cents: 169,
    wasCents: 199,
    category: 'savoury',
    emoji: '🥔',
    tint: '#fff6d4',
    unitPrice: '11,27 € / kg',
    blurb: 'Thick-cut potatoes, fried and salted. Nothing else.',
    tags: ['popular'],
  },
  {
    sku: 'chocolate',
    name: 'Dark chocolate',
    brand: 'Blink',
    pack: '100 g',
    cents: 189,
    category: 'sweet',
    emoji: '🍫',
    tint: '#f6e4d8',
    unitPrice: '18,90 € / kg',
    blurb: '70% cocoa. A clean snap, not too sweet.',
    tags: [],
  },
  {
    sku: 'nuts',
    name: 'Roasted mixed nuts',
    brand: 'Blink',
    pack: '200 g',
    cents: 349,
    category: 'savoury',
    emoji: '🥜',
    tint: '#f8ead4',
    unitPrice: '17,45 € / kg',
    blurb: 'Almonds, cashews, and hazelnuts, lightly salted.',
    tags: [],
  },
  {
    sku: 'pasta-meal',
    name: 'Chicken pesto pasta',
    brand: 'Oven',
    pack: '400 g',
    cents: 549,
    wasCents: 649,
    category: 'meals',
    emoji: '🍝',
    tint: '#eef8e4',
    unitPrice: '13,73 € / kg',
    blurb: 'Heat for four minutes. Dinner, handled.',
    tags: ['popular'],
  },
  {
    sku: 'pizza',
    name: 'Margherita pizza',
    brand: 'Oven',
    pack: '430 g',
    cents: 499,
    category: 'pizza',
    emoji: '🍕',
    tint: '#fff0dc',
    unitPrice: '11,60 € / kg',
    blurb: 'Wood-fired style base, mozzarella, and basil.',
    tags: [],
  },
  {
    sku: 'dumplings',
    name: 'Pea & mint dumplings',
    brand: 'Blink',
    pack: '400 g',
    cents: 399,
    category: 'frozen-meals',
    emoji: '🥟',
    tint: '#e7f8ef',
    unitPrice: '9,98 € / kg',
    blurb: 'From frozen to pan in eight minutes.',
    tags: [],
  },
  {
    sku: 'icecream',
    name: 'Vanilla ice cream',
    brand: 'Daybreak',
    pack: '460 ml',
    cents: 349,
    wasCents: 399,
    category: 'ice-cream',
    emoji: '🍨',
    tint: '#f7f1e6',
    unitPrice: '7,59 € / L',
    blurb: 'Madagascar vanilla, churned with real cream.',
    tags: [],
  },
  {
    sku: 'towel',
    name: 'Kitchen roll',
    brand: 'Blink',
    pack: '4 rolls',
    cents: 229,
    category: 'paper',
    emoji: '🧻',
    tint: '#f4f4f2',
    unitPrice: '0,57 € / roll',
    blurb: 'Strong 2-ply sheets for spills and everything else.',
    tags: [],
  },
  {
    sku: 'soap',
    name: 'Dish soap',
    brand: 'Blink',
    pack: '500 ml',
    cents: 159,
    category: 'cleaning',
    emoji: '🧴',
    tint: '#e5f6f2',
    unitPrice: '3,18 € / L',
    blurb: 'Lemon scent. Cuts grease without a lecture.',
    tags: [],
  },
]

const saved = JSON.parse(localStorage.getItem(CATALOG_KEY)) || {
  categories: seedCategories,
  products: seedProducts,
}

export const categories = reactive(saved.categories)
export const products = reactive(saved.products)

function persist() {
  localStorage.setItem(CATALOG_KEY, JSON.stringify({ categories, products }))
}

export function slugify(text) {
  return text
    .toLowerCase()
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '')
}

export function isCategorySlugTaken(slug) {
  return categories.some(
    (category) =>
      category.slug === slug || category.children.some((child) => child.slug === slug)
  )
}

export function isSkuTaken(sku) {
  return products.some((product) => product.sku === sku)
}

export function categoryUsage(slug) {
  const category = categories.find((item) => item.slug === slug)
  const subcategories = category ? category.children.length : 0
  return products.filter((product) => product.category === slug).length + subcategories
}

export function addCategory({ name, emoji, tint, parent }) {
  const slug = slugify(name)
  const category = { slug, name, emoji, tint }

  if (parent) {
    categories.find((item) => item.slug === parent).children.push(category)
  } else {
    categories.push({ ...category, children: [] })
  }
  persist()
}

export function findCategory(slug) {
  for (const category of categories) {
    if (category.slug === slug) return category

    const child = category.children.find((item) => item.slug === slug)
    if (child) return child
  }
}

export function updateCategory(slug, { name, emoji, tint }) {
  Object.assign(findCategory(slug), { name, emoji, tint })
  persist()
}

export function deleteCategory(slug) {
  const parent = categories.find((item) => item.children.some((child) => child.slug === slug))
  const list = parent ? parent.children : categories
  list.splice(list.findIndex((item) => item.slug === slug), 1)
  persist()
}

export function addProduct(product) {
  products.push(product)
  persist()
}

export function updateProduct(sku, changes) {
  Object.assign(productBySku(sku), changes)
  persist()
}

export function deleteProduct(sku) {
  products.splice(products.findIndex((product) => product.sku === sku), 1)
  persist()
}

export function euros(cents) {
  const amount = (cents / 100).toFixed(2).replace('.', ',')
  return `${amount} €`
}

export function categoryBySlug(slug) {
  for (const category of categories) {
    if (category.slug === slug) return category

    const child = category.children.find((item) => item.slug === slug)
    if (child) return { ...child, parent: category.slug }
  }
}

export function productBySku(sku) {
  return products.find((product) => product.sku === sku)
}

export function productsIn(slug) {
  if (slug === 'deals') return products.filter((product) => product.wasCents)

  const category = categoryBySlug(slug)
  const slugs = category.children?.length ? category.children.map((child) => child.slug) : [slug]

  return products.filter((product) => slugs.includes(product.category))
}

export function searchProducts(query) {
  const needle = query.trim().toLowerCase()
  if (!needle) return []
  return products.filter((product) => {
    const category = categoryBySlug(product.category)
    const parent = category.parent ? categoryBySlug(category.parent) : null
    const haystack =
      `${product.name} ${product.brand} ${product.pack} ${category.name} ${parent?.name || ''}`.toLowerCase()
    return haystack.includes(needle)
  })
}
