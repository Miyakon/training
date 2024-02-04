package caSample

import io.mockk.every
import io.mockk.mockk
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class MenheraKanojoTest {
    @Test
    fun 安定するテスト() {
        val kareshi = mockk<Kareshi>()
        every { kareshi.toMessage() } returns "優しい言葉"
        val target = MenheraKanojo(kareshi)
        assertEquals(target.checkMental(), "安定しました")
    }
}