export const pages = [
  {
    slug: 'about',
    title: 'About',
    lead: 'Groceries from a store nearby, at your door in minutes.',
    sections: [
      {
        heading: 'A store, not a warehouse',
        body: 'Blink takes the shelf from a shop close to you and brings it over in 10–15 minutes. Fruit, milk, dinner, and the kitchen roll you forgot.',
      },
      {
        heading: 'This preview',
        body: 'The catalog, basket, and orders on this site stay on your device. Nothing is charged and no courier is sent. It is here so you can walk the shop before the real one opens.',
      },
    ],
  },
  {
    slug: 'delivery',
    title: 'Delivery',
    lead: 'Pick a window at checkout. We aim for the door, not the depot.',
    sections: [
      {
        heading: 'How fast',
        body: 'As soon as possible is 10–15 minutes. You can also ask for about 30 minutes or about an hour if you are not home yet.',
      },
      {
        heading: 'What it costs',
        body: 'Delivery is 1,90 €. It is free once the basket reaches 29,00 €. The bar in your basket shows how far you have to go.',
      },
      {
        heading: 'Where it goes',
        body: 'The address you save is the door we use. Add a note for the code, the floor, or “leave it by the mat”.',
      },
    ],
  },
  {
    slug: 'help',
    title: 'Help',
    lead: 'The short version of how this shop works.',
    sections: [
      {
        heading: 'Finding something',
        body: 'Open a category, then a subcategory, or use Search. Deals are the products with a pink tag and a crossed-out price.',
      },
      {
        heading: 'Your basket',
        body: 'The plus on a product adds one. Open Basket to change quantities, see delivery, and check out. Your basket stays on this device.',
      },
      {
        heading: 'Orders',
        body: 'After checkout, the order shows under You. It is a preview receipt. No payment is taken.',
      },
      {
        heading: 'A real question',
        body: 'Write to hello@clivern.com and we will get back to you.',
      },
    ],
  },
  {
    slug: 'terms',
    title: 'Terms',
    lead: 'The rules for this preview of the Blink shop.',
    sections: [
      {
        heading: 'A preview, not a sale',
        body: 'Placing an order here does not buy anything. No payment is taken, no store is reserved, and no courier is dispatched.',
      },
      {
        heading: 'The catalog',
        body: 'Prices, packs, and products are sample data. They can change, and they are not an offer from a store.',
      },
      {
        heading: 'Your use',
        body: 'Use the shop as it is shown. Do not rely on an order in this preview as a receipt, a booking, or proof of delivery.',
      },
    ],
  },
  {
    slug: 'privacy',
    title: 'Privacy',
    lead: 'What this shop keeps, and where it keeps it.',
    sections: [
      {
        heading: 'On this device',
        body: 'Your delivery address, basket, and past preview orders are stored in this browser. They are not sent to a Blink account, because this shop does not ask you to sign in.',
      },
      {
        heading: 'What you type at checkout',
        body: 'Name, phone, address, and the courier note are saved with the order on this device so the receipt can show them. Clearing the site data removes them.',
      },
      {
        heading: 'Cookies',
        body: 'The shop preview does not add its own tracking. The rest of Blink asks for cookie choices separately.',
      },
    ],
  },
]

export function pageBySlug(slug) {
  return pages.find((page) => page.slug === slug)
}
