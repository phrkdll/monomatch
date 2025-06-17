<template>
	<article class="relative w-[500px] h-[500px]">
		<div
			v-for="s in placedSymbols"
			:key="s.id"
			class="absolute font-mono whitespace-nowrap"
			:class="clickable ? 'cursor-pointer' : ''"
			:style="{
				top: s.y + 'px',
				left: s.x + 'px',
				fontSize: s.size + 'px',
				transform: `rotate(${s.rotation}deg)`,
				transformOrigin: 'top left',
				color: 'black',
			}"
			@click="() => emit('symbol-click', s.id)"
		>
			{{ s.name }}
		</div>
	</article>
</template>

<script setup>
import { ref, onMounted } from "vue"

const props = defineProps({
	symbols: Array,
	clickable: Boolean,
})
const emit = defineEmits(["symbol-click"])

const placedSymbols = ref([])

const CANVAS_WIDTH = 500
const CANVAS_HEIGHT = 500
const MAX_ATTEMPTS = 500
const MIN_FONT_SIZE = 10
const MAX_FONT_SIZE = 24
const PADDING = 5

function getRotatedBoundingBox(x, y, w, h, angleDeg) {
	const rad = (angleDeg * Math.PI) / 180
	const cos = Math.cos(rad)
	const sin = Math.sin(rad)

	const corners = [
		{ x: 0, y: 0 },
		{ x: w, y: 0 },
		{ x: 0, y: h },
		{ x: w, y: h },
	].map((p) => {
		return {
			x: x + p.x * cos - p.y * sin,
			y: y + p.x * sin + p.y * cos,
		}
	})

	const xs = corners.map(c => c.x)
	const ys = corners.map(c => c.y)

	const minX = Math.min(...xs)
	const minY = Math.min(...ys)
	const maxX = Math.max(...xs)
	const maxY = Math.max(...ys)

	return {
		x: minX,
		y: minY,
		width: maxX - minX,
		height: maxY - minY,
	}
}

function boxesOverlap(a, b) {
	return !(
		a.x + a.width < b.x
		|| a.x > b.x + b.width
		|| a.y + a.height < b.y
		|| a.y > b.y + b.height
	)
}

onMounted(() => {
	const boxes = []

	for (const s of props.symbols) {
		let placed = false
		let fontSize = Math.floor(Math.random() * (MAX_FONT_SIZE - 14)) + 14

		for (let attempt = 0; attempt < MAX_ATTEMPTS; attempt++) {
			const rotation = Math.floor(Math.random() * 181) - 90 // -90 to +90
			const charWidth = fontSize * 0.6
			const w = s.name.length * charWidth
			const h = fontSize

			const maxX = CANVAS_WIDTH - w - PADDING
			const maxY = CANVAS_HEIGHT - h - PADDING

			const x = Math.floor(Math.random() * maxX) + PADDING
			const y = Math.floor(Math.random() * maxY) + PADDING

			const box = getRotatedBoundingBox(x, y, w, h, rotation)

			// Check bounds
			if (
				box.x < 0 || box.y < 0
				|| box.x + box.width > CANVAS_WIDTH
				|| box.y + box.height > CANVAS_HEIGHT
			) {
				continue
			}

			const overlap = boxes.some(b => boxesOverlap(b, box))
			if (!overlap) {
				boxes.push(box)
				placedSymbols.value.push({
					...s,
					x,
					y,
					size: fontSize,
					rotation,
				})
				placed = true
				break
			}

			// If we failed to often, reduce the font size a little bit
			if (attempt > 200 && fontSize > MIN_FONT_SIZE) {
				fontSize -= 1
			}
		}

		if (!placed) {
			console.warn(`Could not place symbol: ${s.name}`)
		}
	}
})
</script>
