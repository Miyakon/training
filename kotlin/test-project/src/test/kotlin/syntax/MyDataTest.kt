package syntax

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Nested
import org.junit.jupiter.api.Test

@Nested
class MyDataTest {

    @Test
    fun getName() {
        val myData = MyData(10)
        assertEquals(myData.test1(), "Hello test")
    }
    
}