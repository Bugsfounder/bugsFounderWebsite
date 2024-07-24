import React from 'react'
import { useParams } from 'react-router'


const Blog = () => {
    let { blog_id } = useParams();

    const blog = {
        id: "1",
        title: "The Future of Web Development",
        content: "Exploring the latest trends and technologies in web development for 2024.",
        author: "John Doe",
        createdAt: "2024-07-01",
        url: "/blogs/The-Future-of-Web-Development"
    }

    return (
        <div className="blogSection dark:text-white p-3">
            <h1 className='class="text-center text-xl md:text-3xl justify-center lg:text-4xl font-semibold text-gray-800 dark:text-white mb-1 mt-[20px] flex"'>{blog.title}</h1>
            <div className='class="flex items-center mb-6 flex-col md:flex-row md:justify-center"'>{blog.author} . {blog.createdAt} . 3 min read</div>
            <p className='class="cont leading-relaxed text-dark dark:text-gray-100"'>{blog.content} Lorem ipsum, dolor sit amet consectetur adipisicing elit. Dolorem ipsum ut iure deserunt alias quo! Veritatis voluptas, recusandae vitae earum quidem itaque velit doloremque quos, corrupti dolores ipsum blanditiis maiores tempore sunt fugiat a neque officia eveniet sapiente exercitationem necessitatibus! Culpa provident doloremque atque totam nam iure optio sed adipisci soluta omnis porro ullam perspiciatis cupiditate quia nisi reprehenderit sequi, ad molestiae. Nam laboriosam ea in consequatur earum veritatis, rerum modi odit, iusto voluptatum, eaque at aliquam ut suscipit numquam molestias amet! Dolore cumque accusamus ab repellendus perspiciatis magnam adipisci dignissimos voluptas fuga beatae similique dolorem ipsa assumenda facilis illum eaque autem, dolores doloremque in laboriosam totam eum consectetur voluptates! Repellat voluptate cum asperiores tempora quos? Ab est, minima doloribus nisi, quod voluptas saepe totam excepturi voluptatibus, mollitia deleniti dolor! Recusandae distinctio dolorem nihil, dignissimos id, minus veritatis omnis voluptatem eaque totam explicabo voluptatibus dolores sed laborum voluptate maxime, nam eius provident fuga ipsam magni voluptates. A tempore suscipit magnam odit veniam! Quidem esse a cum possimus facilis sint, obcaecati nostrum beatae assumenda? Praesentium adipisci modi eveniet sunt voluptates expedita vitae accusamus tempore earum ipsa quia eligendi architecto dolor sequi iste necessitatibus explicabo, sed similique recusandae deleniti, quam animi. Praesentium obcaecati recusandae molestiae, adipisci libero sapiente quia necessitatibus aspernatur beatae numquam fugiat pariatur corrupti sint voluptates. Aliquid officiis magni sunt, sit totam nostrum maxime quas odio beatae excepturi maiores nihil veritatis sed voluptate et vel ea quo atque quia aliquam esse placeat? A, perspiciatis eum quos quam sint error nihil quibusdam magnam? Necessitatibus minus vitae in fugit repellat iusto quaerat similique, temporibus sequi odio sapiente, possimus non ipsam harum nihil repellendus illum. Consectetur doloribus earum optio expedita quasi temporibus, rem non voluptates vel, doloremque excepturi unde! Nostrum sit inventore error, dolor quasi adipisci repellendus pariatur excepturi, tempora earum numquam consequatur facilis debitis. Facilis totam voluptate cupiditate nisi aperiam nobis fugit esse animi qui, mollitia dolorum est eveniet velit hic. Saepe, quidem blanditiis itaque ad omnis voluptatum excepturi est dolore reiciendis accusantium quod laboriosam quasi recusandae vel quas dolores temporibus, dignissimos molestiae architecto maiores aliquid. Culpa beatae iste architecto nihil dolores temporibus odio fuga, ullam nisi facilis recusandae exercitationem, saepe cum numquam labore alias magnam aliquid iusto. Explicabo repudiandae mollitia recusandae ex assumenda non itaque? Deleniti voluptatum aliquam optio modi voluptatibus recusandae minima, maxime provident consectetur repellat id atque doloremque amet voluptates natus explicabo placeat voluptatem suscipit ut consequuntur earum sapiente soluta iure! Deleniti vero consequuntur laudantium, iusto neque possimus id nihil est cupiditate vel esse cumque impedit cum illo nam accusamus ratione asperiores maiores eaque velit repellendus? Asperiores quaerat deleniti sint delectus rerum, porro inventore ullam est itaque, accusantium aliquid dignissimos architecto, temporibus explicabo culpa nostrum quae neque minima amet harum commodi cum quos nisi? Sapiente, sed quaerat magni at praesentium itaque reprehenderit recusandae, iusto aliquid repellendus ut quos repellat nesciunt esse non! Nesciunt beatae porro praesentium, explicabo velit atque mollitia nobis rem laudantium perferendis quae doloribus, error fuga modi! Id repudiandae, esse repellendus ab facilis nemo optio reiciendis sint.</p>
        </div>
    )
}

export default Blog