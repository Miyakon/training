package caSample

class HostKareshi : KareshiInterface{
    override fun sendMessage(): String {
        return listOf<String>("アメ", "ムチ").random()
    }
}
class Kanojo(private val kareshi: KareshiInterface) {
    fun stable(): String {
        return if(kareshi.sendMessage() == "アメ") "安定しました" else "不安定です"
    }
}


