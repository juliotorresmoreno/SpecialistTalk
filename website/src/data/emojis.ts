type Emojis = {
  name: string
  icon: string
  items: Array<[string, string]>
}

const emojis: Emojis[] = [
  {
    name: 'Smileys & People',
    icon: '😃',
    items: require('./_smileys-people.json'),
  },
  {
    name: 'Animals & Nature',
    icon: '🐻',
    items: require('./_animals-nature.json'),
  },
  {
    name: 'Food & Drink',
    icon: '🍔',
    items: require('./_food-drink.json'),
  },
  {
    name: 'Activity',
    icon: '⚽',
    items: require('./_activity.json'),
  },
  {
    name: 'Travel & Places',
    icon: '🚀',
    items: require('./_travel-places.json'),
  },
  {
    name: 'Objects',
    icon: '💡',
    items: require('./_objects.json'),
  },
  {
    name: 'Symbols',
    icon: '💕',
    items: require('./_symbols.json'),
  },
  {
    name: 'Flags',
    icon: '🎌',
    items: require('./_flags.json'),
  },
]

export default emojis
