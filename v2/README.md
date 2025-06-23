# Sunrise Gradient Website

A responsive website showcasing the animated sunrise gradient with interactive controls.

## Files

- `index.html` - Main website file
- `gradient-animation.svg` - 20-second animated gradient (30 stops)
- `mountains-only.png` - Mountain silhouette overlay image

## Features

### Responsive Design
- **Background**: Animated gradient fills entire viewport, can compress horizontally
- **Mask Overlay**: Height matches viewport, width maintains aspect ratio, centered
- **Controls**: Fixed position in upper right corner

### Animation Controls
- **Play Button**: Starts the 20-second animation (appears by default)
- **Pause/Resume**: Pause or resume the current animation
- **Loop Toggle**: Enable/disable looping (green when active)

### Behavior
- Animation does **not** loop by default
- Play button restarts animation from the beginning
- All controls are pure HTML/CSS/JavaScript (no external dependencies)

## Usage

1. Open `index.html` in any modern web browser
2. Click "Play" to start the sunrise animation
3. Use Pause/Resume to control playback
4. Toggle Loop to repeat the animation continuously

## Technical Details

- Pure HTML/CSS/JavaScript implementation
- SVG animation with CSS controls
- Responsive viewport units (vw, vh)
- Backdrop filter effects for controls
- Maintains aspect ratios on all devices

## Browser Compatibility

- Modern browsers with SVG animation support
- CSS backdrop-filter for control styling
- Viewport units (vw, vh) for responsive design